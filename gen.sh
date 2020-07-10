#!/bin/bash
# Script to generate the Protobuf classes from googleapis/google-cloudevents
set -e

echo "~ START"
PROTOBUF_VERSION=3.12.3

# protoc is a native application, so we need to download different zip files
# and use different binaries depending on the OS.
echo "- Determining OS type"
case "$OSTYPE" in
  darwin*)
    PROTOBUF_PLATFORM=osx-x86_64
    PROTOC=tmp/protobuf/bin/protoc
    ;;
  linux*)
    PROTOBUF_PLATFORM=linux-x64_64
    PROTOC=tmp/protobuf/bin/protoc
    ;;
  win* | msys* | cygwin*)
    PROTOBUF_PLATFORM=win64
    PROTOC=tmp/protobuf/bin/protoc.exe
    ;;
  *)
    echo "Unknown OSTYPE: $OSTYPE"
    exit 1
esac

echo "- Cloning github.com/googleapis/google-cloudevents into tmp"
# For the moment, just clone google-cloudevents. Later we might make
# it a submodule. We clone quietly, and only with a depth of 1
# as we don't need history.
rm -rf tmp
mkdir -p tmp
git clone https://github.com/googleapis/google-cloudevents tmp/google-cloudevents -q --depth 1

# We download a specific version rather than using package managers
# for portability and being able to rely on the version being available
# as soon as it's released on GitHub.
echo "- Downloading protobuf tools"
cd tmp
curl -sSL \
  https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/protoc-$PROTOBUF_VERSION-$PROTOBUF_PLATFORM.zip \
  --output protobuf.zip
(mkdir -p protobuf && cd protobuf && unzip -q ../protobuf.zip)
cd ..
chmod +x $PROTOC

echo "- Generating src/ from scratch"

# First delete any previously-generated files if present
rm -rf src 2>/dev/null
mkdir src

# Generate using protoc
$PROTOC \
  --go_out=src/ \
  -I tmp/protobuf/include \
  -I tmp/google-cloudevents/third_party/googleapis \
  -I tmp/google-cloudevents/proto \
  $(find tmp/google-cloudevents/proto -name data.proto)

# Delete temp directory
echo "- Removing tmp/"
rm -rf tmp

echo "~ DONE ✓"
