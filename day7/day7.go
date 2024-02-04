package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

//Enumerating different card values using iota:
const (
	T int = iota + 10
	J
	Q
	K
	A
)
const HandSize int = 5

const (
	HighCard = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	HandType   int
	Cards      string
	CardValues []int //In same order as the individual cards in Cards
}

func (h Hand) String() string {
	return fmt.Sprintln(h.Cards)
}

//Function to read file day7-input.txt
func readFile() []string {
	//call to os.ReadFile returns an array of bytes and an error struct
	data, err := os.ReadFile("./day7-input.txt")
	//Check if there is no error:
	if err != nil {
		fmt.Println("An error occured")
		panic(err)
	}
	//fmt.Print(string(data))
	return strings.Split(string(data), "\n")
}

func processHand(cards string) Hand {

	stringEnumMapping := map[string]int{
		"T": T,
		"J": J,
		"Q": Q,
		"K": K,
		"A": A,
	}
	//Given a hand of cards, parse it and create a Hand struct corresponding to it
	h := Hand{}
	//This map is to determine what is the type of our hand
	cardsMap := make(map[string]int)
	cardsArr := []int{}
	h.Cards = cards
	for _, rune := range cards {
		char := string(rune)
		_, present := cardsMap[char]
		if present {
			cardsMap[char]++
		} else {
			cardsMap[char] = 1
		}
		//Append to cardsArr:
		var intVal int
		if unicode.IsDigit(rune) {
			intVal, _ = strconv.Atoi(char)
		} else {
			intVal, _ = stringEnumMapping[char]
		}
		cardsArr = append(cardsArr, intVal)
	}
	//Set this array to the field CardValues of Hand struct
	h.CardValues = cardsArr
	//Parse through cardsMap and determine
	numOfKeys := len(cardsMap)
	var maxSeen int = 0
	for _, val := range cardsMap {
		if val > maxSeen {
			maxSeen = val
		}
	}
	switch numOfKeys {
	case 1:
		//Five of a kind
		h.HandType = FiveOfAKind
	case 2:
		//Either full house or three of a kind
		if maxSeen == 4 {
			h.HandType = FourOfAKind
		} else {
			h.HandType = FullHouse
		}
	case 3:
		if maxSeen == 3 {
			h.HandType = ThreeOfAKind
		} else {
			h.HandType = TwoPairs
		}
	case 4:
		h.HandType = OnePair
	case 5:
		h.HandType = HighCard
	}
	return h
}

//Get the total using the sorted slice of Hand structs and mapping of Hand to bid value
func getTotal(h []Hand, bids map[string]int) int {
	total := 0
	for ind := 0; ind < len(h); ind++ {
		bidValue, _ := bids[h[ind].Cards]
		total += (ind + 1) * bidValue
	}
	return total
}

func main() {
	//Bids stores the bid value associated with each hand
	bids := make(map[string]int)
	handsArray := []Hand{}
	inp := readFile()
	for ind := 0; ind < len(inp); ind++ {
		lineData := strings.Split(inp[ind], " ")
		//fmt.Println(lineData)
		handString := lineData[0]
		bidValue, _ := strconv.Atoi(lineData[1])
		bids[handString] = bidValue
		handStruct := processHand(handString)
		handsArray = append(handsArray, handStruct)
	}
	//Sort the slice of Hand structs (sort in-place)
	sort.Slice(handsArray, func(i, j int) bool {
		return handsArray[i].HandType < handsArray[j].HandType
	})
	sort.Slice(handsArray, func(i, j int) bool {
		if handsArray[i].HandType == handsArray[j].HandType {
			//Compare values from the CardValues field for the two Hand structs
			for ind := 0; ind < HandSize; ind++ {
				if handsArray[i].CardValues[ind] == handsArray[j].CardValues[ind] {
					continue
				} else {
					return handsArray[i].CardValues[ind] < handsArray[j].CardValues[ind]
				}
			}
		}
		return handsArray[i].HandType < handsArray[j].HandType
	})

	res := getTotal(handsArray, bids)
	fmt.Println("THE ANSWER IS: ", res)
}
