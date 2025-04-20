package main

import (
	"errors"
	"os"
	"testing"
)

/*
# happy path example
$ ./wc -l file.txt
      19 file.txt

# error scenario 1
# note that the file protected_file.txt has no read permission for the user
$ ./wc -l protected_file.txt
./wc: protected_file.txt: open: Permission denied

# error scenario 2
# foo.txt file doesn't exist
$ ./wc -l foo.txt
./wc: foo.txt: open: No such file or directory

# error scenario 3
# bar is a directory, instead of a file
$ ./wc -l bar
./wc: bar: read: Is a directory
*/

// fs package for file system all the interface.
func TestRun(t *testing.T) {
	testCases:= []struct {
		name string
		path string
		want int
		wantErr error
	} {{
		name: "wc over 	non-existent-file",
		path: "new.txt",
		wantErr: os.ErrNotExist,
	},
	{
		name: "wc over 	permission file",
		path: ".",
		wantErr: errors.New("is a directory"),
	},
	{
		name: "file with Lines",
		path: "wc.txt",
		want: 3,
	},
}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := story1(tt.path)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("expected %d lines, got %d", tt.want, got)
			}
		})
	}

}
