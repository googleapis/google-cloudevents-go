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
// 	protoc-gen-go 				 v1.34.1
// source: google/events/cloud/visionai/v1/events.proto

package visionaidata_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/googleapis/google-cloudevents-go/cloud/visionaidata"
	"github.com/googleapis/google-cloudevents-go/internal/testhelper"
	"google.golang.org/protobuf/encoding/protojson"
)

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingAnalysisEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "AnalysisEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.AnalysisEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.AnalysisEventData{}
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
func TestParsingApplicationEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "ApplicationEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.ApplicationEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.ApplicationEventData{}
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
func TestParsingClusterEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "ClusterEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.ClusterEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.ClusterEventData{}
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
func TestParsingDraftEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "DraftEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.DraftEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.DraftEventData{}
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
func TestParsingEventEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "EventEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.EventEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.EventEventData{}
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
func TestParsingProcessEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "ProcessEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.ProcessEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.ProcessEventData{}
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
func TestParsingProcessorEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "ProcessorEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.ProcessorEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.ProcessorEventData{}
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
func TestParsingSeriesEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "SeriesEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.SeriesEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.SeriesEventData{}
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
func TestParsingStreamEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "StreamEventData", "google/events/cloud/visionai/v1")

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
				out := visionaidata.StreamEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := visionaidata.StreamEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}
