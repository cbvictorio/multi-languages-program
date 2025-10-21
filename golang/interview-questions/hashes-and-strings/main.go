package main

import "fmt"

func main() {
	checkTextUniqueness := IsUnique("asa")
	checkStringPermutation := CheckPermutation("ab", "eidboaoo")
	fmt.Printf("\nis unique: %v", checkTextUniqueness)
	fmt.Printf("\ncheck string permutation unique: %v", checkStringPermutation)
}
