package main

import "fmt"

var ROMAN_NUMBERS_MAP = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInteger(text string) int {
	i := 0
	sum := 0

	for i <= len(text)-1 {
		currentChar := string(text[i])
		currentValue := ROMAN_NUMBERS_MAP[currentChar]

		if i == len(text)-1 {
			return sum + currentValue
		}

		nextChar := string(text[i+1])
		nextValue := ROMAN_NUMBERS_MAP[nextChar]

		if currentValue < nextValue {
			sum += nextValue - currentValue
			i += 2
			continue
		}

		sum += currentValue
		i += 1

	}

	return sum
}

func main() {
	romanNumber := "MCMXCIV"
	calculatedRomanNumber := RomanToInteger(romanNumber)
	fmt.Printf("\nThe total calculated roman number of %s is: %d\n", romanNumber, calculatedRomanNumber)
}
