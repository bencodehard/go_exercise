package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := rand.Intn(250)
	fmt.Printf("The value of number is %v\n", number)

	if number >= 0 && number <= 100 {
		fmt.Println("The value of number between 0-100")
	} else if number >= 101 && number <= 200 {
		fmt.Println("The value of number between 101-200")
	} else if number >= 201 && number <= 250 {
		fmt.Println("The value of number between 0-100")
	} else {
		fmt.Printf("The number is %v and over 250", number)
	}
}
