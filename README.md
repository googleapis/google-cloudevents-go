# Google CloudEvents - Go

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/googleapis/google-cloudevents-go)

This library provides classes of common event types used with Google services.

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
