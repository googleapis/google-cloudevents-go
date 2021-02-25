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
package database

import "encoding/json"

// The data within all Firebase Real Time Database reference events.
type ReferenceEventData struct {
	Data  map[string]interface{} `json:"data,omitempty"`  // The original data for the reference.
	Delta map[string]interface{} `json:"delta,omitempty"` // The change in the reference data.
}

func UnmarshalReferenceEventData(data []byte) (ReferenceEventData, error) {
	var d ReferenceEventData
	err := json.Unmarshal(data, &d)
	return d, err
}

func (p *ReferenceEventData) MarshalReferenceEventData() ([]byte, error) {
	return json.Marshal(p)
}
