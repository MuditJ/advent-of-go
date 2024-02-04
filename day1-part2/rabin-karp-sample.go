package main

import (
	"fmt"
	"math"
)

func rabinKarp(pattern, text string) bool {
	//Check if pattern is present in text or not
	//Taking constant for the rolling hash function as 10
	textLen, patternLen := len(text), len(pattern)
	hashValPattern := getHash(pattern)
	fmt.Println("PATTERN IS AND PATTERN HASH VALUE IS: ", pattern, hashValPattern)
	for ind := 0; ind <= (textLen - patternLen); ind++ {
		substring := text[ind : ind+patternLen]
		hashValSubstring := getHash(substring)
		fmt.Println("INDEX IS: ", ind, "AND SUBSTRING IS: ", substring)
		/*
			if ind == 0 {
				hashValSubstring = getHash(substring)
				fmt.Println("NEW SUBSTRING HASH VALUE: ", hashValSubstring)
			} else {
				fmt.Println("Prev hash value was: ", hashValSubstring)
				fmt.Println("Prev char's ascii value is: ", int(substring[0]))
				hashValSubstring = (hashValSubstring-int(substring[0]))*10 + int(substring[patternLen-1])
				fmt.Println("NEW SUBSTRING HASH VALUE: ", hashValSubstring)
			}
		*/
		if hashValSubstring == hashValPattern {
			//Check individual chars, else move forwards
			if substring == pattern {
				fmt.Printf("FOUND MATCH STARTING AT INDEX %d for string %s\n", ind, pattern)
				return true
			}
		} //Else continue
	}
	return false
}

func getHash(s string) int {
	//Get hash value for given string
	base := 256
	mod := int(math.Pow(10, 9)) + 1
	val := 0
	for _, char := range s {
		val = (val*base + int(char)) % mod
	}
	return val
}

func getAsciiSum(s string) int {
	//Return sum of the ascii values for each of the chars in the string passed
	total := 0
	for _, rune := range s {
		fmt.Println(string(rune))
		total += int(rune)
	}
	return total
}
func main() {
	var pattern, text string
	fmt.Println("Enter text to be searched:")
	fmt.Scanln(&text)
	fmt.Println("Enter pattern to be searched:")
	fmt.Scanln(&pattern)
	//fmt.Println(getAsciiSum(text))
	//fmt.Println(getAsciiSum(pattern))

	fmt.Println("Checking for pattern in text..")
	if len(pattern) > len(text) {
		fmt.Printf("PATTERN CANNOT BE LARGER THAN TEXT, NO MATCH FOUND\n")
		return
	}
	res := rabinKarp(pattern, text)
	if res {
		fmt.Println("FOUND PATTERN")
	} else {
		fmt.Println("NOT FOUND")
	}

}
