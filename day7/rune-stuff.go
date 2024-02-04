package main

import "fmt"

const (
	T int = iota + 10
	J
	Q
	K
	A
)

func main() {
	s := "QWERRYT"
	arr := make(map[string]int)
	for _, r := range s {
		char := string(r)
		_, present := arr[char]
		if present {
			arr[char]++
		} else {
			arr[char] = 1
		}
	}
	fmt.Println("ARRAY IS: ")
	fmt.Println(arr)
	fmt.Println((T))
}
