#!/usr/bin/env bash

set -e

SOURCE_OF_TRUTH="../google-cloudevents"

export SNOWPEA_TEST_TEMPLATE_PATH="./generators/protoc-gen-go-snowpea"

_generate() {
    SRC_DIR="${SOURCE_OF_TRUTH}/proto/google/events/"
    proto_src=$(realpath --relative-to="${SRC_DIR}" $1)
    code_dest=$(dirname "${proto_src}")

    # TODO: Build process for audit logs hit issues on proto dependencies.
    # After significant exploration, the solution:
    # Add go_package declarations to relevant protos with package name overrides.
    #  - option go_package = "github.com/googleapis/google-cloudevents-go/pb/cloud/audit/v1;v1";
    #  - option go_package = "github.com/googleapis/google-cloudevents-go/pb/shared/google/apis/monitored_resource;monitored_resource";
    # Use "module" output config.
    if [ "cloud/audit/v1/data.proto" = "${proto_src}" ]; then
        set -x
        protoc \
            -I=../google-cloudevents/proto/google/events/ \
            --go_out=. \
            --go_opt=module=github.com/googleapis/google-cloudevents-go \
            --go-snowpea_out=. \
            --go-snowpea_opt=module=github.com/googleapis/google-cloudevents-go \
            --proto_path=../google-cloudevents/proto \
            --proto_path=../google-cloudevents/third_party/googleapis \
            google/api/monitored_resource.proto cloud/audit/v1/data.proto
        set +x
        return
    fi

    set -x
    # to use snowpea generator, cd generators/protoc-gen-go-snowpea && go install.
    protoc -I=$SRC_DIR --go_out=. \
        --go_opt="M${proto_src}"="pb/${code_dest}" \
        --go-snowpea_out=. \
        --go-snowpea_opt="M${proto_src}"="pb/${code_dest}" \
        --proto_path="${SOURCE_OF_TRUTH}/proto" \
        --proto_path="${SOURCE_OF_TRUTH}/third_party/googleapis" \
        "${proto_src}"
    set +x
}

for i in $(find "${SOURCE_OF_TRUTH}/proto" -type f -name data.proto); do
  _generate $i
done

GOOGLE_CLOUDEVENT_REPO_PATH=$(realpath $SOURCE_OF_TRUTH) go test ./...
