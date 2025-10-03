package main

import (
	"strings"
)

func ConvertStringToArray(str string) []string {
	array := strings.Fields(str)
	return array
}
