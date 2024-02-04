package main

import (
	"fmt"
	"strings"
	"unicode"
)

//Check how Split and FieldsFunc functions are implemented in golang source code
func main() {
	s := "25 32   55 4 1 23435!@@"
	res := strings.FieldsFunc(s, splitFunc)
	fmt.Println(res)
	fmt.Println(len(res)) //6 as there are five spaces and one set of consecutive non-digit chars (!@@) in above string

}

func splitFunc(r rune) bool {
	return unicode.IsDigit(r)
}
