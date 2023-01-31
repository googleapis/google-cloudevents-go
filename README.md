# Google CloudEvents - Go

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/googleapis/google-cloudevents-go) [![Preview](https://img.shields.io/badge/stability-preview-orange?style=flat-square)](https://cloud.google.com/products#section-22)

This library provides Go types for Google CloudEvent data.

## Features

- Simple import and interface
- Inline documentation for Go structs
- Automatic decoding of base64 data
- Enum support
- Protobuf bindings

## Installation

To install this package, run:

``` sh
go get -u github.com/googleapis/google-cloudevents-go
```

This library requires Go 1.17+ and is tested with Go 1.19.

## Usage

Unmarshal a CloudEvent data payload from raw bytes.

<!-- Copied from examples/example_storage_test.go -->
```golang
package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/googleapis/google-cloudevents-go/cloud/storagedata"
	"google.golang.org/protobuf/encoding/protojson"
)

// cloudEventPayload is initialized with an example CloudEvent data payload.
// Source: github.com/googleapis/google-cloudevents/tree/main/examples/binary/storage/StorageObjectData-simple.json
var cloudEventPayload = []byte(`
{
	"bucket": "sample-bucket",
	"contentType": "text/plain",
	"crc32c": "rTVTeQ==",
	"etag": "CNHZkbuF/ugCEAE=",
	"generation": "1587627537231057",
	"id": "sample-bucket/folder/Test.cs/1587627537231057",
	"kind": "storage#object",
	"md5Hash": "kF8MuJ5+CTJxvyhHS1xzRg==",
	"mediaLink": "https://www.googleapis.com/download/storage/v1/b/sample-bucket/o/folder%2FTest.cs?generation=1587627537231057\u0026alt=media",
	"metageneration": "1",
	"name": "folder/Test.cs",
	"selfLink": "https://www.googleapis.com/storage/v1/b/sample-bucket/o/folder/Test.cs",
	"size": "352",
	"storageClass": "MULTI_REGIONAL",
	"timeCreated": "2020-04-23T07:38:57.230Z",
	"timeStorageClassUpdated": "2020-04-23T07:38:57.230Z",
	"updated": "2020-04-23T07:38:57.230Z"
}`)

func Example() {
	data := storagedata.StorageObjectData{}
	if err := protojson.Unmarshal(cloudEventPayload, &data); err != nil {
		log.Fatal("protojson.Unmarshal: ", err)
	}

	updated := data.Updated.AsTime().Format(time.UnixDate)
	fmt.Printf("Bucket: %s, Object: %s, Updated: %s", data.Bucket, data.Name, updated)

	// Output: Bucket: sample-bucket, Object: folder/Test.cs, Updated: Thu Apr 23 07:38:57 UTC 2020
}
```

## Contributing

Contributions to this library are always welcome and highly encouraged.

See [CONTRIBUTING](./CONTRIBUTING.md) for more information how to get started.

Please note that this project is released with a Contributor Code of Conduct.
By participating in this project you agree to abide by its terms.
See Code of Conduct for more information.
