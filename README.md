# Google CloudEvents - Go

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/googleapis/google-cloudevents-go) [![unstable](http://badges.github.io/stability-badges/dist/unstable.svg)](http://github.com/badges/stability-badges)


This library provides classes of common event types used with Google services.

## Installation

**Note**: This library requires Go 1.11+.

To install this package, run:

``` sh
go get -u github.com/googleapis/google-cloudevents-go
```

## Example Usage

Here's an exmaple of using this library with an event from Cloud Pub/Sub with data `MessagePublishedData`.

```go
package main

import (
	"encoding/base64"
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
      "data": "Q2xvdWQgUHViL1N1Yg==",
      "messageId": "136969346945"
    },
    "subscription": "projects/myproject/subscriptions/mysubscription"
  }`)

	var e pubsub.MessagePublishedData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	s, err := base64.URLEncoding.DecodeString(*e.Message.Data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+s\n", s)
}
```

More detailed documentation about the usage of every type can be found in this library's reference.

## Reference

The [`reference.md`](reference.md) file has detailed examples for how to use every event data type.
