package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type FileStat struct {
	lines int
	words int
	chars int
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: wc [-l] [-w] [-c] <filename>")
		os.Exit(1)
	}
	var flags []string
	var filename string

	for _, arg := range args {
		if arg == "-l" || arg == "-w" || arg == "-c" {
			flags = append(flags, arg)
		} else {
			filename = arg
		}
	}

	if filename == "" {
		fmt.Println("Required filename.")
		fmt.Println("Usage: wc [-l] [-w] [-c] <filename>")
		os.Exit(1)
	}

	stat, err := getAllCondition(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(flags) == 0 {
		flags = []string{"-l", "-w", "-c"}
	}

	for _, flag := range flags {
		switch flag {
		case "-l":
			fmt.Printf("%7d ", stat.lines)
		case "-w":
			fmt.Printf("%7d ", stat.words)
		case "-c":
			fmt.Printf("%7d ", stat.chars)
		default:
			fmt.Printf("Unknown flag: %s\n", flag)
			os.Exit(1)
		}
	}
	fmt.Printf("%s\n", filename)
}

func getAllCondition(filepath string) (FileStat, error) {
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
	data, err := validFile(filepath)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(data), "\n")
	return len(lines), nil
}

func countWordInLine(filepath string) (int, error) {
	data, err := validFile(filepath)
	if err != nil {
		return 0, err
	}
	words := strings.Fields(string(data))
	return len(words), nil
}

func countCharsInFile(filepath string) (int, error) {
	data, err := validFile(filepath)
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

func validFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, errors.New("is a directory")
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
