# Google CloudEvents - Go

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/googleapis/google-cloudevents-go) [![unstable](http://badges.github.io/stability-badges/dist/unstable.svg)](http://github.com/badges/stability-badges)

This library provides Go types for Google CloudEvent data.

## Features

- Simple import and interface
- Inline documentation for Go structs
- 64 bit number parsing
- Automatic decoding of base64 data
- Enum support

## Installation

**Note**: This library requires Go 1.11+.

To install this package, run:

``` sh
go get -u github.com/googleapis/google-cloudevents-go
```

## Example Usage

Example event of type `MessagePublishedData` from Cloud Pub/Sub.

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/cloud/pubsub/v1"
)

func main() {
	data := []byte(`{
    "message": {
      "attributes": {
        "key": "value"
      },
      "data": "SGVsbG8sIFdvcmxkIQ==",
      "messageId": "136969346945"
    },
    "subscription": "projects/myproject/subscriptions/mysubscription"
  }`)

	var e pubsub.MessagePublishedData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+s\n", e.Message.Data) // Hello, World!
}
```

More detailed documentation about the usage of every type can be found in this library's reference.

## Reference

The [`reference.md`](reference.md) file has detailed examples for how to use every event data type.
