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

func main() {

	lineflag := flag.Bool("l", false, "Lines")
	wordflag := flag.Bool("w", false, "Words")
	charflag := flag.Bool("c", false, "Characters")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: wc [-l] [-w] [-c] <filenames>")
		os.Exit(1)
	}

	var getflag []string
	if *lineflag {
		getflag = append(getflag, "-l")
	}
	if *wordflag {
		getflag = append(getflag, "-w")
	}
	if *charflag {
		getflag = append(getflag, "-c")
	}

	if len(getflag) == 0 {
		getflag = []string{"-l", "-w", "-c"}
	}

	var wg sync.WaitGroup

	for _, filename := range args {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			all, err := count(file)
			if err != nil {
				fmt.Printf("Error in file %s: %v\n", file, err)
				return
			}
			printCases(all, getflag, file)
		}(filename)
	}

	wg.Wait()
}

func printCases(stat FileStat, flag []string, filename string) {
	for _, cases := range flag {
		switch cases {
		case "-l":
			fmt.Printf("%8d", stat.lines)
		case "-w":
			fmt.Printf("%8d", stat.words)
		case "-c":
			fmt.Printf("%8d", stat.chars)
		default:
			fmt.Printf("Unknown flag: %s\n", flag)
			os.Exit(1)
		}
	}
	fmt.Printf(" %s\n", filename)
}

func count(filepath string) (FileStat, error) {
	countLines, err := countLinesInFile(filepath)
	if err != nil {
		return FileStat{}, err
	}
	countWords, err := countWordInLine(filepath)
	if err != nil {
		return FileStat{}, err
	}
	countChar, err := countCharsInFile(filepath)
	if err != nil {
		return FileStat{}, err
	}

	return FileStat{
		lines: countLines,
		words: countWords,
		chars: countChar,
	}, nil
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
