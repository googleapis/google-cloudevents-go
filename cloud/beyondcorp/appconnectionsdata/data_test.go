// Copyright 2023 Google LLC
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

// Code generated by protoc-gen-go-googlecetypes. DO NOT EDIT.
// versions:
// 	protoc       				 v3.21.6
// 	protoc-gen-go 				 v1.30.0
// 	protoc-gen-go-googlecetypes  short-sha:e11c4a1 (2023-04-20 21:09:22 -0700)
// source: google/events/cloud/beyondcorp/appconnections/v1/events.proto

package appconnectionsdata_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/googleapis/google-cloudevents-go/cloud/beyondcorp/appconnectionsdata"
	"github.com/googleapis/google-cloudevents-go/internal/testhelper"
	"google.golang.org/protobuf/encoding/protojson"
)

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingAppConnectionEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "AppConnectionEventData", "google/events/cloud/beyondcorp/appconnections/v1")

	for name, file := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatal("os.ReadFile:", err)
			}

			if ext := filepath.Ext(file); ext != ".json" {
				t.Fatalf("test support for %q data not implemented", ext)
			}

			t.Run("loose", func(t *testing.T) {
				out := appconnectionsdata.AppConnectionEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := appconnectionsdata.AppConnectionEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}
