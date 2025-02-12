package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Println("Random number 1:", x)
	fmt.Println("Random number 2:", y)

	if x < 4 && y < 4 {
		fmt.Println("Both numbers are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both numbers are greater than 6")
	} else if x >= 4 && x <= 6 && y >= 4 && y <= 6 {
		fmt.Println("The numbers are between 4 and 6")
	} else {
		fmt.Println("The numbers are in different ranges")
	}

}
