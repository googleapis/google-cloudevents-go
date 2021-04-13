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
	fmt.Printf("%+s\n", e.Message.Data)
}
