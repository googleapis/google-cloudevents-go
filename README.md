# Google CloudEvents - Go

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/googleapis/google-cloudevents-go)

This library provides classes of common event types used with Google services.
At this moment the following types are available:

| Package | Struct | Description |
| ------------- | ------------- | ------------- |
| pubsubv1 | MessagePublishedData | A message that is published by publishers and consumed by subscribers. |
| auditv1 | LogEntryData | This event is triggered when a new audit log entry is written. |

## Installation and Usage

**Note**: This library requires Go 1.11+.

To install this package, run

``` sh
go install github.com/googleapis/google-cloudevents-go
```

To use an event class, see the snippet below:

``` go
package main

import (
	"fmt"

	pubsubv1 "github.com/googleapis/google-cloudevents-go/cloud/pubsub/v1"
)

func main() {
	data := []byte("Some event data")
	e, err := pubsubv1.UnmarshalMessagePublishedData(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(e.Message)
}
```
