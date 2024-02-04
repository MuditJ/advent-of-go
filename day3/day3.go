package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Function to read file day3-input.txt

func readFile() [][]string {
	//call to os.ReadFile returns an array of bytes and an error struct
	fileData, err := os.ReadFile("./day3-input.txt")
	//Check if there is no error:
	if err != nil {
		fmt.Println("An error occured while reading the file")
		panic(err)
	}
	//fmt.Print(string(data))
	//Return the input file as an array of strings where each item of array is one line of the file
	var data [][]string
	res := strings.Split(string(fileData), "\n")
	for ind := 0; ind < len(res); ind++ {
		charArray := strings.Split(res[ind], "")
		data = append(data, charArray)
	}
	return data
}

/*
func getPartNumberSum(inputData [][]string) int {
	totalSum := 0
	currentValue := 0
	isPartNumber := false
	cols, rows := len(inputData[0]), len(inputData)
	//Go through each row and for each character, if its a digit, check if in the row above and below if there is an
	//adjacent character which is a special char
	//If so, parse the entire number(sequence of digits)
	//and add it to totalSum
	for rowIndex := 0; rowIndex < rows; rowIndex++ {
		for colIndex := 0; colIndex < cols; colIndex++ {

		}

	}
	return 0
}
*/

func main() {
	/*
		s := "Hello world"
		res := strings.Split(s, "")
		fmt.Println(res)
	*/
	p := "H"
	//r := rune(p)
	//fmt.Println(int(r))
	val, _ := strconv.Atoi(p)
	fmt.Println(val)
	fmt.Println(p)

}
