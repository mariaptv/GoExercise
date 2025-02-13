package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	for x := range x {
		fmt.Println(x)
	}

	m := map[string]int{
		"James": 1,
		"Money": 2,
		"Maria": 3,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
