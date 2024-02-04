package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	RED_MAX   = 12
	GREEN_MAX = 13
	BLUE_MAX  = 14
)

//Function to read file day2-input.txt
func readFile() []string {
	//call to os.ReadFile returns an array of bytes and an error struct
	data, err := os.ReadFile("./day2-input.txt")
	//Check if there is no error:
	if err != nil {
		fmt.Println("An error occured while reading the file")
		panic(err)
	}
	//fmt.Print(string(data))
	//Return the input file as an array of strings where each item of array is one line of the file
	return strings.Split(string(data), "\n")
}

func getGameIdSum(games []string) int {
	total := 0
	totalGames := len(games)
	for ind := 0; ind < totalGames; ind++ {
		allTurnsDetails := strings.Split(games[ind], ":")[1]
		//fmt.Println(allTurnsDetails)
		validGame := true //Game is valid by default
		//Split this further by semicolon and analyze each item in the returned array
		var turns []string = strings.Split(allTurnsDetails, ";")
		//fmt.Println("GAME NUMBER", ind+1, " HAS: ", len(turns), "TURNS")
		//fmt.Println("***********")
		for x := 0; x < len(turns); x++ {
			//fmt.Println(turns[x])
			if !validGame {
				break
			}
			//For each turn in a game, do a further split by "," and check:
			var singleTurnItems []string = strings.Split(turns[x], ",")
			if len(singleTurnItems) == 0 {
				continue
			} else {
				//fmt.Println("+++++++++++++++++++++++++++")
				for y := 0; y < len(singleTurnItems); y++ {
					fmt.Println(singleTurnItems[y])
					item := strings.Split(singleTurnItems[y], " ")
					count, _ := strconv.Atoi(item[1])
					//fmt.Println(item, len(item))
					//fmt.Println(item[1], item[2])
					switch item[2] {
					case "blue":
						if count > BLUE_MAX {
							fmt.Println("BLUE: count is: ", count)
							validGame = false
						}
					case "red":
						if count > RED_MAX {
							fmt.Println("RED: count is: ", count)
							validGame = false
						}
					case "green":
						if count > GREEN_MAX {
							fmt.Println("GREEN: count is: ", count)
							validGame = false
						}

					}
				}
			}

		}
		if validGame {
			fmt.Println("Adding to total an amount of: ", ind+1)
			total += ind + 1
		}
	}
	return total
}

func getPowerCubeSum(games []string) int {
	return 0
}

func main() {
	x := readFile()
	res := getGameIdSum(x)
	fmt.Println(res)

}
