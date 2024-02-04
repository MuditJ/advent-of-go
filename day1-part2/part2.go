package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	Numbers     = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0}
	HashMapping = make(map[int]string)
)

//Read input same way as for part 1, but now
//words need to be parsed to check if there are numbers
//for e.g. tf2th3four -> output should be 2,4

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

//Approach:
//1. Parse input and greedily collect 'characters' into string - apply Rabin Karp Algo to see if it matches with a number
//2. If yes, update either leftmost or rightmost number accordingly
//3. If input character is a digit instead, use same logic as part 1

func getSum(inputArr []string) int {
	//Append 'characters' seen in input greedily to this slice
	totalSum := 0
	for ind := 0; ind < len(inputArr); ind++ {
		fmt.Println("CURRENT: ", inputArr[ind])
		currentWord := []string{}
		leftmostNum, rightmostNum := -1, -1
		for _, rune := range inputArr[ind] {
			if unicode.IsDigit(rune) {
				//Set currentWord to nil
				currentWord = nil
				if leftmostNum == -1 {
					leftmostNum, _ = strconv.Atoi(string(rune))
				}
				rightmostNum, _ = strconv.Atoi(string(rune))
			} else {
				currentWord = append(currentWord, string(rune))
				//Check hash value of this string and see if it is present in the map
				word := strings.Join(currentWord, "")
				currentHash := getHash(word)
				res, found := HashMapping[currentHash]
				if found {
					//Hash value found, compare strings
					if word == res {
						if leftmostNum == -1 {
							leftmostNum, _ = Numbers[word]
						}
						rightmostNum, _ = Numbers[word]
					}
					currentWord = nil
				}
			}
		}
		newVal := (leftmostNum * 10) + rightmostNum
		fmt.Println("NEW VAL IS: ", newVal)
		totalSum += newVal
	}
	return totalSum
}

func main() {
	//Update HashMapping
	for key, _ := range Numbers {
		HashMapping[getHash(key)] = key
	}
	/*
		for key, val := range HashMapping {
			fmt.Println("Key is: ", key, " and val is: ", val)
		}
	*/
	//Read the text file input
	res := readFile()
	fmt.Println(getSum(res))
}
