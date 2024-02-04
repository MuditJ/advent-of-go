package main

import "fmt"

func main() {
	fmt.Println("Checking string lengths for strings containing chars from diff langs")
	s1 := "ABC"
	fmt.Println(len(s1)) //3
	s2 := "$%^"
	fmt.Println(len(s2)) //3
	s3 := "😁😁😁"
	fmt.Println(len(s3)) //12
	s4 := "😁xD"
	fmt.Println(len(s4)) //6

	//We can see that for the smiley emoji, its Unicode code point when it gets
	//converted to its byte representation using UTF-8, takes up 4 bytes

}
