package testhelper

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type dataTestCase struct {
	Ext  string
	Type string
	Case string
}

// FindTestData identifies and loads relevant test cases for a given data type.
func FindTestData(t *testing.T, dataType string, dataPath string) map[string]string {
	t.Helper()
	root := os.Getenv("GENERATE_DATA_SOURCE")
	if root == "" {
		t.Skip("test data: GENERATE_DATA_SOURCE environment variable not set")
	}

	testData, err := AbsDataPath(os.Getenv("GENERATE_DATA_SOURCE"), dataPath)
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

// AbsDataPath normalizes relative or absolute paths for testing.
func AbsDataPath(p, dataPath string) (target string, err error) {
	target = p

	if !filepath.IsAbs(p) {
		target, err = filepath.Abs(p)
		if err != nil {
			return
		}
	}

	// Check if the target path we've assembled to the parent repository exists.
	// This helps identify misconfiguration instead of a lack of tests.
	_, err = os.Stat(target)
	if err != nil {
		return
	}

	// Complete the absolute data path.
	target = filepath.Join(target, "testdata", dataPath)

	return
}
