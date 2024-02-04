package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Function to read file day2-input.txt
func readFile() []string {
	//call to os.ReadFile returns an array of bytes and an error struct
	data, err := os.ReadFile("./day4-input.txt")
	//Check if there is no error:
	if err != nil {
		fmt.Println("An error occured while reading the file")
		panic(err)
	}
	//fmt.Print(string(data))
	//Return the input file as an array of strings where each item of array is one line of the file
	return strings.Split(string(data), "\n")
}

func processInput(s []string) int {
	//Split the input by the vertical separator | or by ':' using the fieldsFunc function
	//from strings package:
	total := 0
	for ind := 0; ind < len(s); ind++ {
		matches := 0
		//This map acts as a set where the keys are the numbers from the winning numbers set
		winningNumsMap := make(map[int]int)
		res := strings.FieldsFunc(s[ind], splitterFunc)
		//Construct map:
		for c := 0; c < 10; c++ {
			num, _ := strconv.Atoi(res[c+2])
			winningNumsMap[num] = 1
		}
		for d := 12; d < len(res); d++ {
			num, _ := strconv.Atoi(res[d])
			//Check if this num is present in the map or not:
			_, present := winningNumsMap[num]
			if present {
				matches++
			}
		}
		if matches > 0 {
			total += int(math.Pow(2, float64(matches-1)))

		}
	}
	return total

}
func splitterFunc(r rune) bool {
	return string(r) == ":" || string(r) == "|" || string(r) == " "
}

func main() {
	fmt.Println("Hello")
	res := processInput(readFile())
	fmt.Println("RESULT IS: ", res)
}
