// Copyright 2021 Google LLC
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

import (
	"encoding/json"
	"testing"
)

func TestMessagePublishedData(t *testing.T) {
	pubsub_data := []byte(`{
    "message": {
      "attributes": {
        "key": "value"
      },
      "data": "SGVsbG8sIFdvcmxkIQ==",
      "messageId": "136969346945"
    },
    "subscription": "projects/myproject/subscriptions/mysubscription"
  }`)

	var e MessagePublishedData
	err := json.Unmarshal(pubsub_data, &e)
	if err != nil {
		panic(err)
	}
	data := string(e.Message.Data)

	if data != "Hello, World!" {
		t.Errorf("Expected %s, got %s", "Hello, World!", data)
	}
}
