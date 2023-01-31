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
// 	protoc-gen-go 				 v1.28.1
// 	protoc-gen-go-googlecetypes  short-sha:906043c (2023-01-19 12:42:43 -0800)
// source: google/events/cloud/iot/v1/events.proto

package iotdata_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/googleapis/google-cloudevents-go/cloud/iotdata"
	"github.com/googleapis/google-cloudevents-go/internal/testhelper"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestMain(m *testing.M) {
	// Tests default the working directory to the test file directory.
	//
	// The process of running the generator then the tests will most likely be
	// done from the module root.
	//
	// This resets the working directory of the tests to the module root so the
	// environment config can be reused across generator & tests.
	err := os.Chdir("../..")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingDeviceEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "DeviceEventData", "google/events/cloud/iot/v1")

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
				out := iotdata.DeviceEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := iotdata.DeviceEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingRegistryEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "RegistryEventData", "google/events/cloud/iot/v1")

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
				out := iotdata.RegistryEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := iotdata.RegistryEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}
