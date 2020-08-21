# Google CloudEvents - Go

This library provides classes of common event types used with Google services.
At this moment the following types are available:

| Package | Struct | Description |
| ------------- | ------------- | ------------- |
| auditv1  | AuditLogWrittenEvent | The event is triggered when a new Cloud Audit Log entry is written. |
| pubsubv1  | MessagePublishedEvent | This event is triggered when a new Cloud Pub/Sub event is published to a topic. |


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
	pubsub "github.com/googleapis/google-cloudevents-go/cloud/pubsub/v1"
)

func main () {
    e := audit.UnmarshalMessagePublishedEvent(data)
    fmt.Println(e.message)
}
```

## Generation

**Note**: Before generating the package,
[quicktype](https://quicktype.io/), and Go 1.11+.

To generate this package, run

``` sh
chmod +x gen.sh
./gen.sh
```