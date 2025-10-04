package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(cwd, "..", "..", "files", "text_file.txt")
	fileContent, err := ReadContentFromFile(filePath)

	if err != nil {
		panic(err)
	}

	fileContentToArray := ConvertTextToArray(fileContent)
	wordsDictionary := CreateWordsDictionary(fileContentToArray)
	wordsLengthMap := CreateWordsLengthMap(fileContentToArray)

	flag := false

	if flag {
		data, _ := json.MarshalIndent(wordsLengthMap, "", "  ")
		fmt.Println(string(data))
	} else {
		for key, value := range wordsDictionary {
			fmt.Printf("(%s) => %d\n", key, value)
		}
	}

}
