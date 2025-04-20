package main

import (
	"errors"
	"os"
	"strings"
)
func story1(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err!= nil {
		return 0, err
	}
	defer file.Close()
	info,_:=file.Stat()
	if info.IsDir(){
		return 0,errors.New("is a directory")
	}
	data, err:= os.ReadFile(filepath)
	if err!=nil {
		return 0,nil
	}
	lines:= strings.Split(string(data), "\n")
	return len(lines), nil
}