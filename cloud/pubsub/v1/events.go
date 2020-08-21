// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    events, err := UnmarshalEvents(bytes)
//    bytes, err = events.Marshal()

package pubsubv1

import "encoding/json"


func UnmarshalMessagePublishedEvent(data []byte) (MessagePublishedEvent, error) {
	var r MessagePublishedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessagePublishedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Events struct {
	MessagePublishedEvent *MessagePublishedEvent `json:"MessagePublishedEvent,omitempty"`// This event is triggered when a new Cloud Pub/Sub event is published to a topic.
}

// This event is triggered when a new Cloud Pub/Sub event is published to a topic.
type MessagePublishedEvent struct {
	Message      *Message `json:"message,omitempty"`     
	Subscription *string  `json:"subscription,omitempty"`
}

type Message struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Message    *string                `json:"message,omitempty"`   
	MessageID  *string                `json:"messageId,omitempty"` 
}
