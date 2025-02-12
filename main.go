package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(250) // Random number between 0 and 99
	fmt.Println("Random number:", randomNumber)
	fmt.Printf("The variable 'randomNumber' is %d\n", randomNumber)
	if randomNumber > 0 && randomNumber < 100 {
		fmt.Println("The random number is between 0 and 100")
	} else if randomNumber > 101 && randomNumber < 200 {
		fmt.Println("The random number is between 101 and 200")
	} else {
		fmt.Println("The random number is between 201 and 250")
	}
}
