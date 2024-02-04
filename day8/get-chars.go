package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "SOMERANDOMSTRING HEREEE"
	byteArr := []byte(s)
	//Get nth character - each character is a single byte
	char, _ := utf8.DecodeRune(byteArr[3:])
	fmt.Println(string(char))

}
