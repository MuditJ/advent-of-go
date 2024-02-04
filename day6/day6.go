package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := solve(readFile())
	fmt.Println("ANSWER IS: ", ans)
	ansTwo := solveTwo(readFile())
	fmt.Println("THE OTHER ANSWER IS: ", ansTwo)
}

//Function to read file day6-input.txt
func readFile() []string {
	//call to os.ReadFile returns an array of bytes and an error struct
	data, err := os.ReadFile("./day6-input.txt")
	//Check if there is no error:
	if err != nil {
		fmt.Println("An error occured")
		panic(err)
	}
	//fmt.Print(string(data))
	return strings.Split(string(data), "\n")
}

func solve(input []string) int {
	result := 1
	times, distances := filterInput(input)
	//fmt.Println(times, len(times))
	//fmt.Println(distances, len(distances))
	//Iterate through the two arrays and get the total number of ways any particular race can be won
	for x := 0; x < len(times); x++ {
		ways := 0
		time, _ := strconv.Atoi(times[x])
		distance, _ := strconv.Atoi(distances[x])
		for y := 1; y < time; y++ {
			if (time-y)*y > distance {
				ways++
			}
		}
		result *= ways
	}
	return result
}

func filterInput(input []string) ([]string, []string) {
	//The first string in this slice is the set of times and the second is the set of distances
	times := strings.FieldsFunc(strings.Join(strings.Split(input[0], "Time:"), " "), filterThis)
	distances := strings.FieldsFunc(strings.Join(strings.Split(input[1], "Distance:"), " "), filterThis)
	return times, distances

}

func solveTwo(input []string) int {
	//Combine all the times into one value and all the distances into one value
	ways := 0
	times, distances := filterInput(input)
	realTime, realDistance := "", ""
	for x := 0; x < len(times); x++ {
		realTime += times[x]
		realDistance += distances[x]
	}
	fmt.Println(realTime, realDistance)
	realTimeInt, _ := strconv.Atoi(realTime)
	realDistInt, _ := strconv.Atoi(realDistance)
	fmt.Println(realDistInt, realTimeInt)
	for y := 1; y < realTimeInt; y++ {
		if (realTimeInt-y)*y > realDistInt {
			ways++
		}
	}
	return ways
}

func filterThis(r rune) bool {
	return string(r) == " "
}
