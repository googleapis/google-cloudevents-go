package testhelper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// PreparePubSubMessagePublishedData modifies data for protojson parsing.
// - Remove fields protojson considers duplicate field names
// - Remove @type properties
func PreparePubSubMessagePublishedData(b []byte) ([]byte, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(b, &j); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	m := j["message"].(map[string]interface{})

	// Remove reserved @type field.
	delete(m, "@type")
	// Remove message_id field name conflict with messageId.
	delete(m, "message_id")
	// Remove publish_time field name conflict with publishTime.
	delete(m, "publish_time")

	j["message"] = m
	return json.Marshal(j)
}

// PrepareAuditLogEntryData modifies data for protojson parsing.
// - Remove @type properties
func PrepareAuditLogEntryData(b []byte) ([]byte, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(b, &j); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	m := j["protoPayload"].(map[string]interface{})

	// Remove reserved @type field.
	delete(m, "@type")

	j["protoPayload"] = m
	return json.Marshal(j)
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
