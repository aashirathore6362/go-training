package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type GrepOptions struct {
	Path   string
	Stdin  io.Reader
	Search string
}

type GrepResult struct {
	MatchLine []string
	Error     error
}

var (
	ErrIsDirectory = errors.New("is a directory")
)

func GrepString(opt GrepOptions) GrepResult {
	var reader io.Reader
	var err error

	if opt.Path != "" {
		reader, err = isValidFile(opt.Path)
		if err != nil {
			return GrepResult{
				Error: err,
			}
		}
		defer reader.(*os.File).Close()
	} else {
		reader = opt.Stdin
	}

	return GrepFromInput(reader, opt.Search)
}

func GrepFromInput(reader io.Reader, keyword string) GrepResult {
	scanner := bufio.NewScanner(reader)
	var matched []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			matched = append(matched, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return GrepResult{
			Error: err,
		}
	}

	return GrepResult{
		MatchLine: matched,
	}
}

func isValidFile(path string) (io.Reader, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, ErrIsDirectory
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
