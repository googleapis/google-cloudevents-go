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
