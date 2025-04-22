package main

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name    string
		path    string
		want    int
		wantErr error
	}{
		{
			name: "Test-1: Count the lines in file",
			path: "testdata/file1.txt",
			want: 5,
		},
		{
			name:    "Test-2: Dir instead of file.",
			path:    ".",
			wantErr: errors.New("is a directory"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countLinesInFile(tt.path)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}

				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Expected error %q but got %q", tt.wantErr, err)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Expected %d but got %d", tt.want, got)
			}
		})
	}
}
func TestCountWord(t *testing.T) {
	testCases := []struct {
		name    string
		path    string
		want    int
		wantErr error
	}{
		{
			name: "Word count for a single file.",
			path: "testdata/file2.txt",
			want: 3,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countWordInLine(tt.path)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}

				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Expected error %q but got %q", tt.wantErr, err)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Expected %d but got %d", tt.want, got)
			}
		})

	}
}
