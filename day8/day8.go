package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pair struct {
	first  string
	second string
}

func getSteps(directions string, nodeMap map[string]Pair) int {
	//Start from AAA and get to ZZZ, and return number of steps taken
	steps := 0
	currentNode := "AAA"
	var nextNode string
	for currentNode != "ZZZ" {
		nextDirection, _ := utf8.DecodeRune([]byte(directions)[steps%len(directions):])
		if string(nextDirection) == "L" {
			nextNode = nodeMap[currentNode].first
		} else {
			nextNode = nodeMap[currentNode].second
		}
		currentNode = nextNode
		steps += 1
	}
	return steps

}

func parseInput(inp []string) (string, map[string]Pair) {
	//Parse the input to return a string representing the directions to take and a map where key
	//is the current node and value is comma separated list of two strings: the left and the right nodes
	nodeMap := make(map[string]Pair)
	directions := inp[0]
	//Parse inp from the third row onwards:
	for row := 3; row < len(inp); row++ {
		data := strings.Split(inp[row], " = ")
		currNode := data[0]
		filteredString := strings.TrimFunc(data[1], func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		//Further split filteredString by the comma:
		nodeString := strings.Split(filteredString, ", ")
		nodeMap[currNode] = Pair{first: nodeString[0], second: nodeString[1]}
	}
	return directions, nodeMap
}

//Function to read file day7-input.txt
func readFile() []string {
	//call to os.ReadFile returns an array of bytes and an error struct
	data, err := os.ReadFile("./day8-input.txt")
	//Check if there is no error:
	if err != nil {
		fmt.Println("An error occured")
		panic(err)
	}
	//fmt.Print(string(data))
	return strings.Split(string(data), "\n")
}

func main() {
	fmt.Println("Hello")
	data := readFile()
	directions, nodeMap := parseInput(data)
	//fmt.Println("DIRECTIONS ARE: ", directions)
	//fmt.Println("MAP IS: ", nodeMap)
	res := getSteps(directions, nodeMap)
	fmt.Println("ANSWER IS: ", res)
}
