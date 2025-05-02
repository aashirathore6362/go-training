package main

import (
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
	// options := wordCountOptions{}
	lineflag := flag.Bool("l", false, "Count lines")
	wordflag := flag.Bool("w", false, "Count words")
	charflag := flag.Bool("c", false, "Count characters")

	flag.Parse()
	args := flag.Args()

	options := optionCount(*lineflag, *wordflag, *charflag)

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

func optionCount(lineflag, wordflag, charflag bool) wordCountOptions {
	if !lineflag && !wordflag && !charflag {
		// No flags given, enable all
		return wordCountOptions{
			isLineCount: true,
			isWordCount: true,
			isCharCount: true,
		}
	}
	return wordCountOptions{
		isLineCount: lineflag,
		isWordCount: wordflag,
		isCharCount: charflag,
	}
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
	data, err := readFile(filepath)
	if err != nil {
		return FileStat{}, err
	}
	// files := string(data)
	if options.isLineCount {
		stats.lines, _ = countLinesInFile(data)
	}
	if options.isWordCount {
		stats.words, _ = countWordInLine(data)
	}
	if options.isCharCount {
		stats.chars, _ = countCharsInFile(data)
	}

	return stats, nil
}
func countLinesInFile(data []byte) (int, error) {
	// data, err := readFile(filepath)
	// if err != nil {
	// 	return 0, err
	// }
	lines := strings.Split(string(data), "\n")
	return len(lines), nil
}

func countWordInLine(data []byte) (int, error) {
	// data, err := readFile(filepath)
	// if err != nil {
	// 	return 0, err
	// }
	words := strings.Fields(string(data))
	return len(words), nil
}

func countCharsInFile(data []byte) (int, error) {
	// data, err := readFile(filepath)
	// if err != nil {
	// 	return 0, err
	// }
	charCount := 0
	words := strings.Fields(string(data))
	for _, word := range words {
		charCount += len(word)
	}
	if len(words) > 0 {
		return charCount + len(words) - 1, nil
	}
	return 0, nil
}

func validateFile(filepath string) error {
	info, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("%s is a directory", filepath)
	}
	return nil
}
func readFileData(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func readFile(filepath string) ([]byte, error) {
	if err := validateFile(filepath); err != nil {
		return nil, err
	}
	return readFileData(filepath)
	// data, err := os.ReadFile(filepath)
	// if err != nil {
	// 	return nil, err
	// }
	// return data, nil
}
