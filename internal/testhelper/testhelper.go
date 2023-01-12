package testhelper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// CleanAcceptedUnknowns strips test data of undefined fields that we know will
// be present in data payloads. These are exceptions for strict validation testing.
// These are primarily '@type' fields injected as an artifact of protobuf.Unknown
// being serialized into the data upstream from the event.
func CleanAcceptedUnknowns(b []byte, dataType string, typePrefix string) ([]byte, error) {
	switch {
	case typePrefix == "google.events.cloud.pubsub.v1" && dataType == "MessagePublishedData":
		var j map[string]interface{}
		if err := json.Unmarshal(b, &j); err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}
		m := j["message"].(map[string]interface{})
		delete(m, "@type")
		j["message"] = m
		fmt.Println("Removed [message.@type] from test data.")
		return json.Marshal(j)
	case typePrefix == "google.events.cloud.audit.v1" && dataType == "LogEntryData":
		var j map[string]interface{}
		if err := json.Unmarshal(b, &j); err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}
		m := j["protoPayload"].(map[string]interface{})
		delete(m, "@type")
		j["protoPayload"] = m
		fmt.Println("Removed [protoPayload.@type] from test data.")
		return json.Marshal(j)
	}

	return b, nil
}

type dataTestCase struct {
	Ext  string
	Type string
	Case string
}

// FindTestData identifies and loads relevant test cases for a given data type.
func FindTestData(t *testing.T, dataType string, dataPath string) map[string]string {
	t.Helper()
	testDataRoot := os.Getenv("GENERATE_DATA_SOURCE")
	if testDataRoot == "" {
		t.Skip("test data: GENERATE_DATA_SOURCE environment variable not set")
	}

	// Tests are run from the working directory of the test file.
	// Tests using this function are assumed to be nested two directories down
	// from the repository root.
	testData := filepath.Join("../..", testDataRoot, "testdata", dataPath)
	testData, err := filepath.Abs(testData)
	if err != nil {
		t.Fatal(err)
	}

	files, err := os.ReadDir(testData)
	if err != nil {
		t.Skip("No test cases found: os.ReadDir:", err)
	}

	cases := make(map[string]string, len(files))
	for _, file := range files {
		c := dataTestCase{
			Ext: filepath.Ext(file.Name()),
		}
		if c.Ext != ".json" {
			continue
		}

		var ok bool
		c.Type, c.Case, ok = strings.Cut(strings.TrimSuffix(file.Name(), c.Ext), "-")
		if !ok {
			c.Case = "default"
		}

		if c.Type == dataType {
			cases[c.Case] = filepath.Join(testData, file.Name())
		}
	}

	if len(cases) == 0 {
		t.Skip("No test cases found for", dataType)
	}

	return cases
}
