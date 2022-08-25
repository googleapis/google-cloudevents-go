#!/usr/bin/env bash

set -e

SOURCE_OF_TRUTH="../google-cloudevents"

_generate() {
    SRC_DIR="${SOURCE_OF_TRUTH}/proto/google/events/"
    proto_src=$(realpath --relative-to="${SRC_DIR}" $1)
    code_dest=$(dirname "${proto_src}")

    # TODO: Build process breaks on audit logs:
    #     protoc-gen-go: unable to determine Go import path for "google/api/monitored_resource.proto"
    # Please specify either:
    # 	• a "go_package" option in the .proto source file, or
    # 	• a "M" argument on the command line.
    # See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.
    # --go_out: protoc-gen-go: Plugin failed with status code 1.
    if [ "cloud/audit/v1/data.proto" = "${proto_src}" ]; then
        return
    fi

    echo
    echo "generating ${proto_src}..."
    set -x
    protoc -I=$SRC_DIR --go_out=. \
        --go_opt="M${proto_src}"="pb/${code_dest}" \
        --proto_path="${SOURCE_OF_TRUTH}/proto" \
        --proto_path="${SOURCE_OF_TRUTH}/third_party/googleapis" \
        "${proto_src}"
    set +x
}
for i in $(find "${SOURCE_OF_TRUTH}/proto" -type f -name data.proto); do
  _generate $i
done