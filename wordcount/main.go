package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

type FileStat struct {
	lines int
	words int
	chars int
}

type wordCountOptions struct {
	isLineCount bool
	isWordCount bool
	isCharCount bool
}

func main() {
	// var options wordCountOptions
	options := wordCountOptions{}
	lineflag := flag.Bool("l", false, "Count lines")
	wordflag := flag.Bool("w", false, "Count words")
	charflag := flag.Bool("c", false, "Count characters")

	flag.Parse()
	args := flag.Args()

	if !*lineflag && !*wordflag && !*charflag {
		options = wordCountOptions{
			isLineCount: true,
			isWordCount: true,
			isCharCount: true,
		}
	} else {
		options = wordCountOptions{
			isLineCount: *lineflag,
			isWordCount: *wordflag,
			isCharCount: *charflag,
		}
	}

	var wg sync.WaitGroup

	for _, filename := range args {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			stats, err := count(file, options)
			if err != nil {
				fmt.Printf("Error in file %s: %v\n", file, err)
				return
			}
			printStats(stats, options, file)
		}(filename)
	}

	wg.Wait()
}

func printStats(stat FileStat, options wordCountOptions, filename string) {
	if options.isLineCount {
		fmt.Printf("%8d", stat.lines)
	}
	if options.isWordCount {
		fmt.Printf("%8d", stat.words)
	}
	if options.isCharCount {
		fmt.Printf("%8d", stat.chars)
	}
	fmt.Printf(" %s\n", filename)
}

func count(filepath string, options wordCountOptions) (FileStat, error) {
	var stats FileStat

	if options.isLineCount {
		stats.lines, _ = countLinesInFile(filepath)
	}
	if options.isWordCount {
		stats.words, _ = countWordInLine(filepath)
	}
	if options.isCharCount {
		stats.chars, _ = countCharsInFile(filepath)
	}

	return stats, nil
}
func countLinesInFile(filepath string) (int, error) {
	data, err := readFile(filepath)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(data), "\n")
	return len(lines), nil
}

func countWordInLine(filepath string) (int, error) {
	data, err := readFile(filepath)
	if err != nil {
		return 0, err
	}
	words := strings.Fields(string(data))
	return len(words), nil
}

func countCharsInFile(filepath string) (int, error) {
	data, err := readFile(filepath)
	if err != nil {
		return 0, err
	}
	charCount := 0
	words := strings.Fields(string(data))
	for _, word := range words {
		charCount += len(word)
	}
	return charCount + len(words) - 1, nil
}

func validateFile(filepath string) error {
	info, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("is a directory")
	}
	return nil
}

func readFile(filepath string) ([]byte, error) {
	if err := validateFile(filepath); err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
