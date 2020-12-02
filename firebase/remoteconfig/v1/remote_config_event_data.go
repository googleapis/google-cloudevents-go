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
package remoteconfig

import "encoding/json"

// The data within all Firebase Remote Config events.
type RemoteConfigEventData struct {
	Description    *string     `json:"description,omitempty"`   // The user-provided description of the corresponding Remote Config template.
	RollbackSource *string     `json:"rollbackSource,omitempty"`// Only present if this version is the result of a rollback, and will be the; version number of the Remote Config template that was rolled-back to.
	UpdateOrigin   *Update     `json:"updateOrigin"`            // Where the update action originated.
	UpdateTime     *string     `json:"updateTime,omitempty"`    // When the Remote Config template was written to the Remote Config server.
	UpdateType     *Update     `json:"updateType"`              // What type of update was made.
	UpdateUser     *UpdateUser `json:"updateUser,omitempty"`    // Aggregation of all metadata fields about the account that performed the update.
	VersionNumber  *string     `json:"versionNumber,omitempty"` // The version number of the version's corresponding Remote Config template.
}

// Aggregation of all metadata fields about the account that performed the update.
type UpdateUser struct {
	Email    *string `json:"email,omitempty"`   // Email address.
	ImageURL *string `json:"imageUrl,omitempty"`// Image URL.
	Name     *string `json:"name,omitempty"`    // Display name.
}

type Update struct {
	Integer *int64
	String  *string
}

func UnmarshalRemoteConfigEventData(data []byte) (RemoteConfigEventData, error) {
	var d RemoteConfigEventData
	err := json.Unmarshal(data, &d)
	return d, err
}

func (p *RemoteConfigEventData) MarshalRemoteConfigEventData() ([]byte, error) {
	return json.Marshal(p)
}