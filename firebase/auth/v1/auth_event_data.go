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

package auth

// AuthEventData: The data within all Firebase Auth events.
type AuthEventData struct {
	CustomClaims  map[string]interface{} `json:"customClaims,omitempty"`  // User's custom claims, typically used to define user roles and propagated; to an authenticated user's ID token.
	Disabled      *bool                  `json:"disabled,omitempty"`      // Whether the user is disabled.
	DisplayName   *string                `json:"displayName,omitempty"`   // The user's display name.
	Email         *string                `json:"email,omitempty"`         // The user's primary email, if set.
	EmailVerified *bool                  `json:"emailVerified,omitempty"` // Whether or not the user's primary email is verified.
	Metadata      *Metadata              `json:"metadata,omitempty"`      // Additional metadata about the user.
	PhoneNumber   *string                `json:"phoneNumber,omitempty"`   // The user's phone number.
	PhotoURL      *string                `json:"photoURL,omitempty"`      // The user's photo URL.
	ProviderData  []UserInfo             `json:"providerData,omitempty"`  // User's info at the providers
	Uid           *string                `json:"uid,omitempty"`           // The user identifier in the Firebase app.
}

// Metadata: Additional metadata about the user.
type Metadata struct {
	CreateTime     *string `json:"createTime,omitempty"`     // The date the user was created.
	LastSignInTime *string `json:"lastSignInTime,omitempty"` // The date the user last signed in.
}

// UserInfo: User's info at the identity provider
type UserInfo struct {
	DisplayName *string `json:"displayName,omitempty"` // The display name for the linked provider.
	Email       *string `json:"email,omitempty"`       // The email for the linked provider.
	PhotoURL    *string `json:"photoURL,omitempty"`    // The photo URL for the linked provider.
	ProviderID  *string `json:"providerId,omitempty"`  // The linked provider ID (e.g. "google.com" for the Google provider).
	Uid         *string `json:"uid,omitempty"`         // The user identifier for the linked provider.
}
