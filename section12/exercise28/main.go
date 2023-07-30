package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("The value  of x is %v and y is %v\n", x, y)

	switch {
	case x < 4 && y < 4:
		{
			fmt.Println("both are less than four")
		}
	case x > 6 && y > 6:
		{
			fmt.Println("both are more than six")
		}
	case x >= 4 && x <= 6:
		{
			fmt.Println("x is between 4 to 6")
		}
	case y != 5:
		{
			fmt.Println("y is not five")
		}
	default:
		{
			fmt.Println("none of the previous are met")
		}

	}
}
