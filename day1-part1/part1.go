package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var totalSum int = 0

//Function to read file day1-input.txt
func readFile() []string {
	//call to os.ReadFile returns an array of bytes and an error struct
	data, err := os.ReadFile("./day1-input.txt")
	//Check if there is no error:
	if err != nil {
		fmt.Println("An error occured")
		panic(err)
	}
	//fmt.Print(string(data))
	return strings.Split(string(data), "\n")
}

func getSum(inputArr []string) int {

	//Approach: The entire file has been read into a single string
	//Read the input using for range loop, stopping at each rune representing a newline char
	//Get the leftmost seen integer and the rightmost seen integer (these might be same)
	//Calculate leftmostInt * 10 + rightmostInt
	//Set both to -1 and continue
	totalSum := 0
	for ind := 0; ind < len(inputArr); ind++ {
		leftmostNum, rightmostNum := -1, -1
		for _, rune := range inputArr[ind] {
			if unicode.IsDigit(rune) {
				if leftmostNum == -1 {
					leftmostNum, _ = strconv.Atoi(string(rune))
				}
				rightmostNum, _ = strconv.Atoi(string(rune))
			}
		}
		totalSum += (leftmostNum * 10) + rightmostNum
	}
	return totalSum
}

func main() {
	//Reading the file present in current directory
	res := readFile()
	fmt.Println(getSum(res))
}
