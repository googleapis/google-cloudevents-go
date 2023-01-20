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

# Setup environment for code generation:
# - install protobuf tools in a local temp directory
# - clone google-cloudevents repo if needed
set -e

name=$(basename ${BASH_SOURCE:-$0})
library_version=$(git rev-parse --short HEAD)
library_date=$(git show -s --format=%ci "${library_version}")
echo "google-cloudevents-go > ${name} (${library_version} on ${library_date})"
echo "working-directory: ${PWD}"
echo

if [ "${GENERATE_DATA_SOURCE:0:3}" = "tmp" ] || [ "${GENERATE_DATA_SOURCE:0:5}" = "./tmp" ]; then
  if [ -d "${GENERATE_DATA_SOURCE}" ]; then
    echo "setup-generator.sh will delete all contents of ./tmp. Currently GENERATE_DATA_SOURCE='${GENERATE_DATA_SOURCE}'."
    exit 2
  fi
fi

# Create a location for local tool installation.
echo "- Removing existing tmp/ directory"
rm -rf tmp
mkdir tmp

# Protobuf Setup
PROTOBUF_VERSION=21.6
# protoc is a native application, so we need to download different zip files
# and use different binaries depending on the OS.
echo "- Determining OS type"
case "$OSTYPE" in
  linux*)
    PROTOBUF_PLATFORM=linux-x86_64
    PROTOC=tmp/protobuf/bin/protoc
    ;;
  win* | msys* | cygwin*)
    PROTOBUF_PLATFORM=win64
    PROTOC=tmp/protobuf/bin/protoc.exe
    ;;
  darwin*)
    PROTOBUF_PLATFORM=osx-x86_64
    PROTOC=tmp/protobuf/bin/protoc
    ;;    
  *)
    echo "Unknown OSTYPE: $OSTYPE"
    exit 1
esac

# We download a specific version rather than using package managers
# for portability and being able to rely on the version being available
# as soon as it's released on GitHub.
echo "- Downloading protobuf tools..."
pushd tmp 1>/dev/null
curl -sSL \
  https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/protoc-$PROTOBUF_VERSION-$PROTOBUF_PLATFORM.zip \
  --output protobuf.zip
(mkdir protobuf && cd protobuf && unzip -q ../protobuf.zip)
popd 1>/dev/null
chmod +x $PROTOC

echo "- Downloaded protobuf ${PROTOBUF_VERSION} for ${PROTOBUF_PLATFORM}"

echo "- Downloading & installing the Go protocol buffers plugin..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
echo "- Protobuf tooling installation complete"

echo "- Installing custom generators"
pushd generators/protoc-gen-go-googlecetypes 1>/dev/null
go install .
popd 1>/dev/null

if [[ -z "${GENERATE_DATA_SOURCE}" ]]; then
  echo "- Cloning github.com/googleapis/google-cloudevents into tmp"
  # For the moment, just clone google-cloudevents. Later we might make
  # it a submodule. We clone quietly, and only with a depth of 1
  # as we don't need history.
  dest='tmp/google-cloudevents'
  git clone https://github.com/googleapis/google-cloudevents "${dest}" -q --depth 1
else
  echo "- Skipping new clone of github.com/googleapis/google-cloudevents, GENERATE_DATA_SOURCE variable already set"
fi

echo
echo "Configure environment for generate_code.sh:"
echo "- Usage: Configure the path to protobuf tools ('export GENERATE_PROTOC_PATH=$PROTOC')"
if [[ -z "${GENERATE_DATA_SOURCE}" ]]; then
  echo "- Usage: Configure the path to proto definitions ('export GENERATE_DATA_SOURCE=\$(realpath ${dest})')"
fi

echo
echo "You are ready to run 'sh generate-code.sh'"
