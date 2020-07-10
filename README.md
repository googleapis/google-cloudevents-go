# Google CloudEvents - Go

This repository contains types for CloudEvents issued by Google.

## Prerequisites

- Go 1.9 or higher.
- Go protobuf tool
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go`

## Generate

Generate libraries:

```sh
./gen.sh
```

Observe the new files in `src/`.