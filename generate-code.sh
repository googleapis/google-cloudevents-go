#!/usr/bin/env bash
# Copyright 2022, Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Generates code from https://github.com/googleapis/google-cloudevents:
# - Types from protobuf messages
#
# Configuration:
# - GENERATE_DATA_SOURCE: Path to google-cloudevent repo.
# - GENERATE_PROTOC_PATH: Path to protobuf tool. 
#
# Usage:
# - sh ./build.sh
# - GENERATE_DATA_SOURCE=tmp/google-cloudevents GENERATE_PROTOC_PATH=protoc sh ./build.sh

set -e

# Output Utilities
_heading() {
  echo
  echo "$(tput bold)${1}$(tput sgr0)"
}

name=$(basename "${BASH_SOURCE[0]}")
library_version=$(git rev-parse --short HEAD)
library_date=$(git show -s --format=%ci "${library_version}")
echo "google-cloudevents-go > ${name} (${library_version} on ${library_date})"

# Required configuration.
if [[ -z "${GENERATE_DATA_SOURCE}" ]]; then
  echo
  echo "Environment variable 'GENERATE_DATE_SOURCE' not found."
  echo "Please run 'sh tools/setup-generator.sh' and follow the instructions."
  exit 1
fi

if [[ -z "${GENERATE_PROTOC_PATH}" ]]; then
  echo
  echo "Environment variable 'GENERATE_PROTOC_PATH' not found."
  echo "Please run 'sh tools/setup_generator.sh' and follow the instructions."
  exit 1
fi

# Derive proto repo metadata.
data_version=$(git -C "${GENERATE_DATA_SOURCE}" rev-parse --short HEAD)
data_date=$(git -C "${GENERATE_DATA_SOURCE}" show -s --format=%ci "${data_version}")
# Derive proto lookup paths.
src_dir="${GENERATE_DATA_SOURCE}/proto/google/events"
googleapis_dir="${GENERATE_DATA_SOURCE}/third_party/googleapis"
# Prepare dependencies for build & generation.
monitored_proto="google/api/monitored_resource.proto"

_heading "Preparing to generate library..."
echo "- Schema Source Repository: \t${GENERATE_DATA_SOURCE} (${data_version} on ${data_date})"
echo "- Proto Source Directory: \t${src_dir}"
echo "- Shared googleapis Protos: \t${googleapis_dir}"

# Manifest file with details about how code was most recently generated.
# Useful when troubleshooting without build logs.
cat > generated.txt <<GENERATION_METADATA
created_date: $(date "+%Y-%m-%d %T %z")
data_commit_date: ${data_date}
data_commit_hash: ${data_version}
tool_commit_date: ${library_date}
tool_commit_hash: ${library_version}

GENERATION_METADATA

# Clean up previously generated files.
# - Prevent continued presence of files we no longer generate
# - Troubleshooting assist to see what failed to generate
rm -rf cloud firebase shared

# Generate proto code that is a dependency for an event type.
# Module mappings for generation are different from module mappings
# for import needed when configuring the dependents.
_heading "Generating dependencies..."
echo "- ${monitored_proto}"

$GENERATE_PROTOC_PATH --go_out=. \
  --go_opt="M${monitored_proto}"="shared/google;google" \
  --proto_path="${googleapis_dir}" \
  "$googleapis_dir/${monitored_proto}"

_generateData() {
    proto_src=$(realpath --relative-to="${src_dir}" "$1")
    code_dest=$(dirname "$(dirname "${proto_src}")")
    product=$(basename "$code_dest")
    version=$(basename "$(dirname "${proto_src}")")

    # Explicit type versioning after v1.
    if [[ "${version}" == "v1" ]]; then
      version=""
    fi

    echo "- ${product}: ${proto_src} => ${code_dest}data${version}"

    $GENERATE_PROTOC_PATH --go_out=. \
      --go_opt="M${proto_src}"="${code_dest}data${version};${product}data${version}" \
      --go_opt="M${monitored_proto}"="github.com/googleapis/google-cloudevents-go/shared/google;google" \
      --proto_path="${src_dir}" \
      --proto_path="${googleapis_dir}" \
      "${proto_src}"
}

_heading "Generating data type code in Go..."
for i in $(find "${GENERATE_DATA_SOURCE}/proto" -type f -name data.proto); do
  _generateData "$i"
done
