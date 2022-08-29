# Google CloudEvents Go

This is a prototype rebuild of this library using protobuf as the source of
truth in place of the JSON-Schema definition.

## Usage

This prototype generates a couple type helper functions.

```golang
import (
    "log"
    cloudevent "github.com/cloudevents/sdk-go/v2"
    "github.com/googleapis/google-cloudevents-go/pb/cloud/pubsub/v1"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Does not exist: https://github.com/cloudevents/sdk-go/issues/766
    event := cloudevent.EventFromRequest(r)
    data, err := event.DataBytes()
    if err != nil {
        log.Fatal("event.DataBytes:", err)
    }
    
    obj, err := pubsub.UnmarshalMessagePublishedData(event.DataContentType(), data)
    if err != nil {
        log.Fatal("pubsub.UnmarshalMessagePublishedData:", err)
    }
    log.Print(obj.Message.Data)
}
```

## Repository Map

```txt
.
├── README.md
├── build.sh: Generate code for all types & run tests
├── generators
│   └── protoc-gen-go-snowpea: Generate helper functions and tests.
│       ├── generator.go
│       ├── go.mod
│       ├── go.sum
│       └── validationtest.gotpl
├── go.mod
├── go.sum
├── pb: Generated code (protobufs, tests, helpers)
└── testing
```

## Testing

For usage details, review build.sh.

* [x] Parse testdata from google-cloudevents repository, ignoring unknown fields
* [ ] Parse testdata from google-cloudevents repository, strict.
* [ ] Test helper functions

## Open Questions

How to do strict validation?
