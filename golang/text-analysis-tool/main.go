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
	wordsLengthMap := CreateWordsLengthMap(fileContentToArray)

	data, _ := json.MarshalIndent(wordsLengthMap, "", "  ")
	fmt.Println(string(data))

}
