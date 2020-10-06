// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    messagePublishedData, err := UnmarshalMessagePublishedData(bytes)
//    bytes, err = messagePublishedData.Marshal()

package pubsubv1

import "encoding/json"

func UnmarshalMessagePublishedData(data []byte) (MessagePublishedData, error) {
	var r MessagePublishedData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessagePublishedData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// A message that is published by publishers and consumed by subscribers.
type MessagePublishedData struct {
	Message      *PubsubMessage `json:"message,omitempty"`     // The message that was published.
	Subscription *string        `json:"subscription,omitempty"`// The resource name of the subscription for which this event was generated. The format of; the value is `projects/{project-id}/subscriptions/{subscription-id}`.
}

// The message that was published.
//
// A message published to a topic.
type PubsubMessage struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`// Attributes for this message. If this field is empty, the message must contain non-empty; data. This can be used to filter messages on the subscription.
	Data       *string                `json:"data,omitempty"`      // The message data field. If this field is empty, the message must contain at least one; attribute. A base64-encoded string.
	MessageID  *string                `json:"messageId,omitempty"` // ID of this message, assigned by the server when the message is published. Guaranteed to; be unique within the topic. This value may be read by a subscriber that receives a; PubsubMessage via a subscriptions.pull call or a push delivery. It must not be populated; by the publisher in a topics.publish call.
}
