package testhelper_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/googleapis/google-cloudevents-go/internal/testhelper"
)

func TestAbsDataPath(t *testing.T) {
	// The directory of this test file.
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Testing WD:", wd)

	cases := []struct {
		label     string
		inputBase string
		inputPath string
		want      string
		wantError bool
	}{
		{
			label:     "empty",
			inputBase: "",
			inputPath: "",
			want:      filepath.Join(wd, "testdata"),
		},
		{
			label:     "relative",
			inputBase: "",
			inputPath: "exists.json",
			want:      filepath.Join(wd, "testdata/exists.json"),
		},
		{
			label:     "absolute",
			inputBase: wd,
			inputPath: "exists.json",
			want:      filepath.Join(wd, "testdata/exists.json"),
		},
		{
			label:     "wrong",
			inputBase: wd + "404",
			inputPath: "exists.json",
			// Error is returned before inputPath join.
			want:      filepath.Join(wd + "404"),
			wantError: true,
		},
		{
			label:     "no-product",
			inputBase: wd,
			inputPath: "404.json",
			want:      filepath.Join(wd, "testdata/404.json"),
			wantError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			got, err := testhelper.AbsDataPath(tc.inputBase, tc.inputPath)
			if tc.wantError && err == nil {
				t.Error("want err, got nil")
			} else if !tc.wantError && err != nil {
				t.Error(err)
			}
			if got != tc.want {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}
