package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func getFirstAndLastNumbers(input string) (first, last int, err error) {
	var firstDigit, lastDigitFound bool

	for _, char := range input {
		if unicode.IsDigit(char) {
			if !firstDigit {
				// Convert the first digit found to an integer
				first, err = strconv.Atoi(string(char))
				if err != nil {
					return 0, 0, err
				}
				firstDigit = true
			}

			// Convert the current digit to an integer and update the last digit
			last, err = strconv.Atoi(string(char))
			if err != nil {
				return 0, 0, err
			}

			lastDigitFound = true
		}
	}

	if !lastDigitFound {
		return 0, 0, fmt.Errorf("no digits found in the input string")
	}

	return first, last, nil
}

func main() {
	input := "abc123def456"

	first, last, err := getFirstAndLastNumbers(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("First digit: %d\nLast digit: %d\n", first, last)
}
