package main

import "fmt"

func main() {
	var s string = "😁😁😁"
	//Iterating though this string
	fmt.Println(len(s)) //12
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	//Output is:
	//f0 9f 98 81 f0 9f 98 81 f0 9f 98 81 % -> 12 bytes in total as shown by above line's output
	//ITERATING A STRING ACCESSES THE INDIVIDUAL BYTES - NOT THE INDIVIDUAL CHARACTERS
}
