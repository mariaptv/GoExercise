package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 100 {
		if i%2 == 0 {
			fmt.Println("number ", i)
		}
		i++
	}
}
