
# Under Heavy Development

This library is not ready for use. We are working to make it ready. For more details, see #113.

You can access the previous development version of this library at
[v0.2.3](https://github.com/googleapis/google-cloudevents-go/tree/v0.2.3)

# Google CloudEvents - Go

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/googleapis/google-cloudevents-go) [![unstable](http://badges.github.io/stability-badges/dist/unstable.svg)](http://github.com/badges/stability-badges)

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

### Unmarshal Data

Unmarshal the CloudEvent data payload from raw bytes.

```golang
import (
	"log"
	"time"

	"github.com/googleapis/google-cloudevents-go/cloud/storagedata"
	"google.golang.org/protobuf/encoding/protojson"
)

func UnmarshalData(cloudEventPayload []byte) {
	data := storagedata.StorageObjectData{}
	if err := protojson.Unmarshal(cloudEventPayload, &data); err != nil {
		log.Fatal("protojson.Unmarshal: ", err)
	}

	updated := data.Updated.AsTime().Format(time.UnixDate)
	log.Printf("Bucket: %s, Object: %s, Updated: %s", data.Bucket, data.Name, updated)
}
```

### Unmarshal Data from HTTP Request

Unmarshal the CloudEvent Data payload from an HTTP request containing
a "binary-mode" CloudEvent.

```go
import (
	"log"
	"time"

	cloudevent "github.com/cloudevents/sdk-go/v2"
	"github.com/googleapis/google-cloudevents-go/cloud/storagedata"
	"google.golang.org/protobuf/encoding/protojson"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	event, err := cloudevent.NewEventFromHTTPRequest(r)
	if err != nil {
		log.Fatal("cloudevent.NewEventFromHTTPRequest:", err)
	}
	cloudEventPayload := event.Data()

	data := storagedata.StorageObjectData{}
	if err := protojson.Unmarshal(cloudEventPayload, &data); err != nil {
		log.Fatal("protojson.Unmarshal: ", err)
	}

	updated := data.Updated.AsTime().Format(time.UnixDate)
	log.Printf("Bucket: %s, Object: %s, Updated: %s", data.Bucket, data.Name, updated)
}
```
