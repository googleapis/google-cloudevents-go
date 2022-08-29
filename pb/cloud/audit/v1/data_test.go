// Code generated by protoc-gen-go-snowpea. DO NOT EDIT.
package v1

import (
	"path/filepath"
	"os"
	"strings"
	"testing"
)

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - TODO: "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingLogEntryData(t *testing.T) {
	cases := findTestData(t, "LogEntryData")

	for name, file := range cases {
		data, err := os.ReadFile(file)
		if err != nil {
			t.Fatal("os.ReadFile:", err)
		}

		obj, err := UnmarshalLogEntryData("application/json", data)
		if err != nil {
			t.Errorf("%s: parsing failed: %v", name, err)
		} else {
			t.Log(name, obj.String())
		}
	}
}

const (
	DATA_TYPE = 0
	EXTENSION = 1
	TEST_CASE = 1
)

func findTestData(t *testing.T, dataType string) map[string]string {
	t.Helper()
	testDataRoot := os.Getenv("GOOGLE_CLOUDEVENT_REPO_PATH")
	if testDataRoot == "" {
		t.Skip("test data: GOOGLE_CLOUDEVENT_REPO_PATH environment variable not set")
	}

	testData := filepath.Join(testDataRoot, "testdata", "google", "events", "googleapis/google-cloudevents-go/pb/cloud/audit/v1")
	files, err := os.ReadDir(testData)
	if err != nil {
		t.Skip("No test cases found: ioutil.ReadDir:", err)
	}

	cases := make(map[string]string, len(files))
	for _, file := range files {
		parts := strings.Split(file.Name(), ".")
		if parts[EXTENSION] != "json" {
			continue
		}
		metadata := strings.Split(parts[0], "-")
		if metadata[DATA_TYPE] != dataType {
			continue
		}
		cases[metadata[TEST_CASE]] = filepath.Join(testData, file.Name())
	}

	if len(cases) == 0 {
		t.Skip("No test cases found for", dataType)
	}

	return cases
}
