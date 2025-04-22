package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Usage: binary -l <filename>")
		os.Exit(1)
	}
	option := args[1]
	filepath := args[2]
	switch option {
	case "-l":
		lines, err := countLinesInFile(filepath)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Lines:", lines)
	case "-w":
		wordCount, err := countWordInLine(filepath)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Words:", wordCount)
	default:
		fmt.Println("Unknown option:", option)
		fmt.Println("Use -l for lines, -w word count")
		os.Exit(1)
	}
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
	wordCount := strings.Fields(string(data))
	return len(wordCount), err
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
