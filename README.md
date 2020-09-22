# Google CloudEvents - Go

This library provides classes of common event types used with Google services.
At this moment the following types are available:

| Package | Struct | Description |
| ------------- | ------------- | ------------- |
| cloud.pubsub.v1 | MessagePublishedEvent | This event is triggered when a Pub/Sub message is published. |
| cloud.audit.v1 | AuditLogWrittenEvent | This event is triggered when a new audit log entry is written. |

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
	e, err := pubsubv1.UnmarshalMessagePublishedEvent(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(e.Message)
}
```
