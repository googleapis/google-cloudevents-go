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
# - GENERATE_TEMPLATE_DIR: Location of gotemplate files for custom code generation.
#
# Usage:
# - sh ./generate-code.sh
# - GENERATE_DATA_SOURCE=tmp/google-cloudevents GENERATE_PROTOC_PATH=protoc sh ./generate-code.sh

set -e

# Output Utilities
_heading() {
  if [ -t 1 ];  then
    echo
    echo "$(tput bold)${1}$(tput sgr0)"
  else
    echo "=> ${1}"
  fi
}

name=$(basename "${BASH_SOURCE[0]}")
library_version=$(git rev-parse --short HEAD)
library_date=$(git show -s --format=%ci "${library_version}")
echo "google-cloudevents-go > ${name} (${library_version} on ${library_date})"
echo "working-directory: ${PWD}"

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

if [[ -z "${GENERATE_TEMPLATE_DIR}" ]]; then
  export GENERATE_TEMPLATE_DIR=$(realpath ./generators/templates)
fi

GENERATE_DATA_SOURCE=$(realpath "${GENERATE_DATA_SOURCE}")

# Derive proto repo metadata.
data_version=$(git -C "${GENERATE_DATA_SOURCE}" rev-parse --short HEAD)
data_date=$(git -C "${GENERATE_DATA_SOURCE}" show -s --format=%ci "${data_version}")
# Derive proto lookup paths.
src_dir="${GENERATE_DATA_SOURCE}/proto/google/events"
googleapis_dir="${GENERATE_DATA_SOURCE}/third_party/googleapis"

_heading "Preparing to generate library..."
echo "- Schema Source Repository: \t${GENERATE_DATA_SOURCE} (${data_version} on ${data_date})"
echo "- Proto Source Directory: \t${src_dir}"
echo "- Shared googleapis Protos: \t${googleapis_dir}"
echo "- Custom Code Generator Templates: \t${GENERATE_TEMPLATE_DIR}"

# Manifest file with details about how code was most recently generated.
# Useful when troubleshooting without build logs.
cat > generated.txt <<GENERATION_METADATA
data_commit_date: ${data_date}
data_commit_hash: ${data_version}
tool_commit_date: ${library_date}
tool_commit_hash: ${library_version}

GENERATION_METADATA

export PROTOC_GEN_GO_VERSION="$(protoc-gen-go --version | awk  '{print $2}')"
export LIBRARY_VERSION="short-sha:${library_version} (${library_date})"


# Clean up previously generated files.
# - Prevent continued presence of files we no longer generate
# - Troubleshooting assist to see what failed to generate
rm -rf cloud firebase shared

# Generate proto code that is a dependency for an event type.
# Module mappings for generation are different from module mappings
# for import needed when configuring the dependents.
_heading "Generating dependencies..."
deps=$(find "${GENERATE_DATA_SOURCE}/third_party/googleapis/google/api" -type f -name *.proto)
echo $deps |  xargs realpath --relative-to="${GENERATE_DATA_SOURCE}" | xargs printf -- '- %s\n'

#/third_party/googleapis \
$GENERATE_PROTOC_PATH --go_out=. \
   --go_opt=module=github.com/googleapis/google-cloudevents-go \
   --proto_path="${googleapis_dir}" \
   "${deps}"

_generateData() {
    proto_src=$(realpath --relative-to="${src_dir}" "$1")
    code_dest=$(dirname "$(dirname "${proto_src}")")
    product=$(basename "$code_dest")
    version=$(basename "$(dirname "${proto_src}")")

    # Explicit type versioning after v1.
    if [[ "${version}" == "v1" ]]; then
      version=""
    fi

    # Skip unsupported products.
    if [[ "${product}" == "pubsub" ]]; then
      echo "- ${product}: skipping generation, not currently supported"
      return
    fi

    echo "- ${product}: ${proto_src} => ${code_dest}data${version}"

    $GENERATE_PROTOC_PATH --go_out=. \
      --go_opt="M${proto_src}"="${code_dest}data${version};${product}data${version}" \
      --proto_path="${src_dir}" \
      --proto_path="${googleapis_dir}" \
      "${proto_src}"
}

_generateValidationTests() {
    # Source the schema protos.
    proto_dir="${GENERATE_DATA_SOURCE}/proto"
    proto_src=$(realpath --relative-to="${proto_dir}" "$1")
    # Derive path to data.proto from event.proto.
    # This is consistently available but not currently a standard.
    data_src=$(dirname "${proto_src}")/data.proto
    # Derive destination directory.
    code_dest=$(dirname "$(dirname "${proto_src}")" | cut -d'/' -f3-)
    # Product name & API version
    product=$(basename "$code_dest")
    version=$(basename "$(dirname "${proto_src}")")

    # Explicit type versioning after v1.
    if [[ "${version}" == "v1" ]]; then
      version=""
    fi

    # Skip unsupported products.
    if [[ "${product}" == "pubsub" ]]; then
      echo "- ${product}: skipping generation, not currently supported"
      return
    fi

    echo "- ${product}: ${proto_src} => ${code_dest}data${version}/data_test.go"

    $GENERATE_PROTOC_PATH --go-googlecetypes_out=. \
        --go-googlecetypes_opt="Mgoogle/events/cloudevent.proto"="github.com/googleapis/google-cloudevents-go/thirdparty/cloudevents;cloudevents" \
        --go-googlecetypes_opt="M${proto_src}"="${code_dest}data${version}" \
        --go-googlecetypes_opt="M${data_src}"="${code_dest}data${version};${product}data${version}" \
        --proto_path="${proto_dir}" \
        --proto_path="${googleapis_dir}" \
        "${proto_src}"
}

_heading "Generating data type code in Go..."
for i in $(find "${GENERATE_DATA_SOURCE}/proto" -type f -name data.proto); do
  _generateData "$i"
done

_heading "Generating validation tests..."
for i in $(find "${GENERATE_DATA_SOURCE}/proto" -type f -name events.proto); do
  _generateValidationTests "$i"
done

_heading "Running validation tests..."
go test -v ./...

_heading "Running example tests..."
pushd examples >/dev/null
go test -v ./...
popd >/dev/null
