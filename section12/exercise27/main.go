package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("value of x is %v and y is %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("both are less than four")
	} else if x > 6 && y > 6 {
		fmt.Println("both are more than six")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is between 4 to 6")
	} else if y != 5 {
		fmt.Println("y is not five")
	} else {
		fmt.Println("none of the previous are met")
	}
}
