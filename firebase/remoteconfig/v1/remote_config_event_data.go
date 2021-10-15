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

package remoteconfig

// RemoteConfigEventData: The data within all Firebase Remote Config events.
type RemoteConfigEventData struct {
	Description    *string       `json:"description,omitempty"`           // The user-provided description of the corresponding Remote Config template.
	RollbackSource *int64        `json:"rollbackSource,string,omitempty"` // Only present if this version is the result of a rollback, and will be the; version number of the Remote Config template that was rolled-back to.
	UpdateOrigin   *UpdateOrigin `json:"updateOrigin,omitempty"`          // Where the update action originated.
	UpdateTime     *string       `json:"updateTime,omitempty"`            // When the Remote Config template was written to the Remote Config server.
	UpdateType     *UpdateType   `json:"updateType,omitempty"`            // What type of update was made.
	UpdateUser     *User         `json:"updateUser,omitempty"`            // Aggregation of all metadata fields about the account that performed the; update.
	VersionNumber  *int64        `json:"versionNumber,string,omitempty"`  // The version number of the version's corresponding Remote Config template.
}

// User: Aggregation of all metadata fields about the account that performed the
// update.
//
// All the fields associated with the person/service account
// that wrote a Remote Config template.
type User struct {
	Email    *string `json:"email,omitempty"`    // Email address.
	ImageURL *string `json:"imageUrl,omitempty"` // Image URL.
	Name     *string `json:"name,omitempty"`     // Display name.
}

// Where the update action originated.
type UpdateOrigin string

const (
	AdminSDKNode                        UpdateOrigin = "ADMIN_SDK_NODE"
	Console                             UpdateOrigin = "CONSOLE"
	RESTAPI                             UpdateOrigin = "REST_API"
	RemoteConfigUpdateOriginUnspecified UpdateOrigin = "REMOTE_CONFIG_UPDATE_ORIGIN_UNSPECIFIED"
)

// What type of update was made.
type UpdateType string

const (
	ForcedUpdate                      UpdateType = "FORCED_UPDATE"
	IncrementalUpdate                 UpdateType = "INCREMENTAL_UPDATE"
	RemoteConfigUpdateTypeUnspecified UpdateType = "REMOTE_CONFIG_UPDATE_TYPE_UNSPECIFIED"
	Rollback                          UpdateType = "ROLLBACK"
)
