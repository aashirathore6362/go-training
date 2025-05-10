package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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
	lineflag := flag.Bool("l", false, "Count lines")
	wordflag := flag.Bool("w", false, "Count words")
	charflag := flag.Bool("c", false, "Count characters")
	flag.Parse()

	args := flag.Args()
	options := optionCount(*lineflag, *wordflag, *charflag)

	if len(args) == 0 {
		stats, err := countFromReader(os.Stdin, options)
		if err != nil {
			fmt.Printf("Error reading stdin: %v\n", err)
			return
		}
		printData(stats, options, "")
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		stats, err := countFromReader(file, options)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filename, err)
			continue
		}
		printData(stats, options, filename)
	}
}

func countFromReader(r io.Reader, options wordCountOptions) (FileStat, error) {
	var stats FileStat

	data, err := io.ReadAll(r)
	if err != nil {
		return stats, err
	}

	if options.isLineCount {
		stats.lines = countLines(data)
	}
	if options.isWordCount {
		stats.words = countWords(data)
	}
	if options.isCharCount {
		stats.chars = len(data)
	}
	return stats, nil
}

func countLines(data []byte) int {
	return len(strings.Split(string(data), "\n"))
}

func countWords(data []byte) int {
	return len(strings.Fields(string(data)))
}

func optionCount(lineflag, wordflag, charflag bool) wordCountOptions {
	if !lineflag && !wordflag && !charflag {
		return wordCountOptions{true, true, true}
	}
	return wordCountOptions{
		isLineCount: lineflag,
		isWordCount: wordflag,
		isCharCount: charflag,
	}
}

func printData(stat FileStat, options wordCountOptions, filename string) {
	if options.isLineCount {
		fmt.Printf("%8d", stat.lines)
	}
	if options.isWordCount {
		fmt.Printf("%8d", stat.words)
	}
	if options.isCharCount {
		fmt.Printf("%8d", stat.chars)
	}
	if filename != "" {
		fmt.Printf(" %s", filename)
	}
	fmt.Println()
}
