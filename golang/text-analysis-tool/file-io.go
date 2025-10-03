package main

import (
	"errors"
	"os"
)

func ReadContentFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fileContent := string(data)

	if fileContent == "" {
		return "", errors.New("the file is empty")
	}

	return fileContent, nil
}
