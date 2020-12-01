# Google CloudEvents - Go

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/googleapis/google-cloudevents-go)

This library provides classes of common event types used with Google services.

## Installation

**Note**: This library requires Go 1.11+.

To install this package, run

``` sh
go install github.com/googleapis/google-cloudevents-go
```

## Example Usage

Here's an exmaple of using this library with an event from Cloud Pub/Sub with data `MessagePublishedData`.

```go
package main

import (
	"fmt"

	pubsub "github.com/googleapis/google-cloudevents-go/cloud/pubsub/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	e, err := pubsub.UnmarshalMessagePublishedData(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

More detailed documentation about the usage of every type can be found in this library's reference.
## Reference

The [`reference.md`](reference.md) file has detailed examples for how to use every event data type.
