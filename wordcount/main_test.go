package main

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	testdatas := []struct {
		name    string
		options wordCountOptions
		path    string
		want    FileStat
		wantErr error
	}{
		{
			name:    "Dir instead of file.",
			path:    "testdata",
			wantErr: errors.New("is a directory"),
		},
		{
			name:    "File does not exist",
			path:    "testdata/missing.txt",
			wantErr: errors.New("no such file"),
		},
		{
			name: "Test lines, words, chars count for a single file",
			path: "testdata/file1.txt",
			options: wordCountOptions{
				isLineCount: true,
				isWordCount: true,
				isCharCount: true,
			},
			want: FileStat{
				lines: 5,
				words: 5,
				chars: 21,
			},
		},
		{
			name: "Test lines, words, chars count for a single file",
			path: "testdata/file1.txt",
			options: wordCountOptions{
				isLineCount: false,
				isWordCount: true,
				isCharCount: true,
			},
			want: FileStat{
				words: 5,
				chars: 21,
			},
		},
		{
			name: "wc -l with single match",
			path: "testdata/file2.txt",
			options: wordCountOptions{
				isLineCount: true,
			},
			want: FileStat{
				lines: 1,
			},
		},
		{
			name: "wc -w with no matches",
			path: "testdata/file2.txt",
			options: wordCountOptions{
				isWordCount: true,
			},
			want: FileStat{
				words: 0,
			},
		},
		{
			name: "wc -c with matches",
			path: "testdata/file1.txt",
			options: wordCountOptions{
				isCharCount: true,
			},
			want: FileStat{
				chars: 21,
			},
		},
		{
			name: "wc -c with multiple matches",
			path: "testdata/file3.txt",
			options: wordCountOptions{
				isCharCount: true,
			},
			want: FileStat{
				chars: 68,
			},
		},
		{
			name: "wc -lc with multiple matches",
			path: "testdata/file3.txt",
			options: wordCountOptions{
				isLineCount: true,
				isCharCount: true,
			},
			want: FileStat{
				lines: 8,
				chars: 68,
			},
		},
		{
			name: "wc -wc with multiple matches",
			path: "testdata/file3.txt",
			options: wordCountOptions{
				isWordCount: true,
				isCharCount: true,
			},
			want: FileStat{
				words: 12,
				chars: 68,
			},
		},
		{
			name: "wc -lw with multiple matches",
			path: "testdata/file3.txt",
			options: wordCountOptions{
				isLineCount: true,
				isWordCount: true,
			},
			want: FileStat{
				lines: 8,
				words: 12,
			},
		},
		{
			name: "wc -lwc with multiple matches",
			path: "testdata/file3.txt",
			options: wordCountOptions{
				isLineCount: true,
				isWordCount: true,
				isCharCount: true,
			},
			want: FileStat{
				lines: 8,
				words: 12,
				chars: 68,
			},
		},
		{
			name: "wc -lwc with multiple matches",
			path: "testdata/file3.txt",
			options: wordCountOptions{
				isLineCount: true,
				isWordCount: true,
				isCharCount: true,
			},
			want: FileStat{
				lines: 8,
				words: 12,
				chars: 68,
			},
		},
	}
	for _, tt := range testdatas {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.path)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("Unexpected error opening file: %v", err)
				}
				if !strings.Contains(err.Error(), tt.wantErr.Error()) {
					t.Fatalf("Expected error to contain %q but got %q", tt.wantErr.Error(), err.Error())
				}
				return
			}
			defer f.Close()

			got, err := countFromReader(f, tt.options)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("Expected error but got nil")
				}
				if !strings.Contains(err.Error(), tt.wantErr.Error()) {
					t.Fatalf("Expected error to contain %q but got %q", tt.wantErr.Error(), err.Error())
				}

				return
			}
			if err != nil {
				t.Fatalf("Unexpected error while reading file: %v", err)
			}
			if got.lines != tt.want.lines {
				t.Errorf("Expected %d lines but got %d", tt.want.lines, got.lines)
			}
			if got.words != tt.want.words {
				t.Errorf("Expected %d words but got %d", tt.want.words, got.words)
			}
			if got.chars != tt.want.chars {
				t.Errorf("Expected %d chars but got %d", tt.want.chars, got.chars)
			}
		})
	}
}
func TestFromStdin(t *testing.T) {
	stdin := []byte("one\ntwo three\nfour five six")

	options := wordCountOptions{
		isLineCount: true,
		isWordCount: true,
		isCharCount: true,
	}

	reader := strings.NewReader(string(stdin))

	got, err := countFromReader(reader, options)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := FileStat{
		lines: 3,
		words: 6,
		chars: len(stdin),
	}

	if got.lines != want.lines {
		t.Errorf("expected %d lines, got %d", want.lines, got.lines)
	}
	if got.words != want.words {
		t.Errorf("expected %d words, got %d", want.words, got.words)
	}
	if got.chars != want.chars {
		t.Errorf("expected %d chars, got %d", want.chars, got.chars)
	}
}
