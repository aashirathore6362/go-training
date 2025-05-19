package main

import (
	"bytes"
	"errors"
	"slices"
	"testing"
)

func TestMatchString(t *testing.T) {
	testCases := []struct {
		name     string
		stdin    []byte
		fileName string
		Search   string
		result   GrepResult
		expErr   error
	}{
		{
			name:     "greps a single-line file",
			fileName: "testdata/file1.txt",
			Search:   "is",
			result:   GrepResult{MatchLine: []string{"is"}},
			expErr:   nil,
		},
		{
			name:     "Check is dir",
			fileName: "testdata",
			expErr:   ErrIsDirectory,
		},
		{
			name:   "read std",
			stdin:  []byte("this\nis\nfile\ngo"),
			Search: "go",
			result: GrepResult{MatchLine: []string{"go"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			options := GrepOptions{
				Path:   tc.fileName,
				Stdin:  bytes.NewReader(tc.stdin),
				Search: tc.Search,
			}
			got := GrepString(options)
			want := tc.result

			if tc.expErr != nil {
				if got.Error == nil {
					t.Fatalf("Expected an error but didn't get one")
				}
				if !errors.Is(got.Error, tc.expErr) {
					t.Fatalf("Expected error %q but got %q", tc.expErr.Error(), got.Error.Error())
				}
				return
			}
			if got.Error != nil {
				t.Fatalf("Didn't expect an error: %v", got.Error)
			}
			if !slices.Equal(got.MatchLine, want.MatchLine) {
				t.Errorf("Expected %v but got %v", want.MatchLine, got.MatchLine)
			}
		})
	}
}
