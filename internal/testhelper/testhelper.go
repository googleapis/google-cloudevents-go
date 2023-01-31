package testhelper

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type dataTestCase struct {
	// Ext is the test data file extension, used to inform data format.
	Ext string
	// Type is the CloudEvent data type inferred from the testdata filename.
	Type string
	// Case is the testing use case inferred from the testdata filename.
	Case string
}

// FindTestData identifies and loads relevant test cases for a given data type.
func FindTestData(t *testing.T, dataType string, dataPath string) map[string]string {
	t.Helper()
	root := os.Getenv("GENERATE_DATA_SOURCE")
	if root == "" {
		t.Fatal("test data: GENERATE_DATA_SOURCE environment variable not set")
	}

	_, err := os.Stat(root)
	if err != nil {
		t.Fatal("test data: GENERATE_DATA_SOURCE environment variable not set to an existing base path.\nSet this variable to the local path of the google-cloudevents repository.")
	}

	if !filepath.IsAbs(root) {
		t.Fatal("test data: GENERATE_DATA_SOURCE environment variable not an absolute path")
	}

	testData := filepath.Join(root, "testdata", dataPath)
	files, err := os.ReadDir(testData)
	if err != nil {
		t.Skip("No test cases found: os.ReadDir:", err)
	}

	cases := make(map[string]string, len(files))
	for _, file := range files {
		c := dataTestCase{
			Ext: filepath.Ext(file.Name()),
		}
		if c.Ext != ".json" && c.Ext != ".proto" {
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
