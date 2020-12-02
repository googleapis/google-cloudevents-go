// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package pubsub

import "encoding/json"

// The event data when a message is published to a topic.
type MessagePublishedData struct {
	Message      *Message `json:"message,omitempty"`     // The message that was published.
	Subscription *string  `json:"subscription,omitempty"`// The resource name of the subscription for which this event was; generated. The format of the value is; `projects/{project-id}/subscriptions/{subscription-id}`.
}

// The message that was published.
type Message struct {
	Attributes map[string]string `json:"attributes,omitempty"`// Attributes for this message.
	Data       *string           `json:"data,omitempty"`      // The binary data in the message.
	MessageID  *string           `json:"messageId,omitempty"` // ID of this message, assigned by the server when the message is published.; Guaranteed to be unique within the topic.
}

func UnmarshalMessagePublishedData(data []byte) (MessagePublishedData, error) {
	var d MessagePublishedData
	err := json.Unmarshal(data, &d)
	return d, err
}

func (p *MessagePublishedData) MarshalMessagePublishedData() ([]byte, error) {
	return json.Marshal(p)
}