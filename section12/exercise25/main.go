package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(250)
	fmt.Printf("the value of x is %v\n", x)

	switch {
	case x <= 100:
		{
			fmt.Println("x is less than 100")
		}

	case x >= 101 && x <= 200:
		{
			fmt.Println("x is between 101 and 200")
		}

	case x >= 201 && x <= 250:
		{
			fmt.Println("x is between 201 and 250")
		}

	default:
		{
			fmt.Println("x is over 250")
		}
	}
}
