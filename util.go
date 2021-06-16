package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Walk(path string) []string {
	var out []string
	err := filepath.Walk(path, func(cp string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println("Entering Directory:", cp)
		} else {
			out = append(out, cp)
		}
		return nil
	})
	if err != nil {
		panic(errors.New("error at scanning"))
	}
	return out
}

func Split(path string) []string {
	var content []string
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(errors.New("read file failed"))
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content
}

func DerefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
