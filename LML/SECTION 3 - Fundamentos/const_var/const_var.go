package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.14159
	var radius = 3.2

	// shorter version
	area := PI * math.Pow(radius, 2)
	fmt.Println("The area is", area);

	const (
		a = 1
		b = 2
		c = 3
	)

	var (	
		d = 4
		e = 5
	)

	fmt.Println(a, b, c, d, e)

	var f, g bool = true, false
	fmt.Println(f, g)

	// Arithmetic operations
	fmt.Println("6 * 4 =",  6 * 4)
	fmt.Println("6 / 4 =",  6 / 4)
	fmt.Println("6 + 4 =",  6 + 4)
	fmt.Println("6 - 4 =",  6 - 4)
	fmt.Println("6 %% 4 =", 6 % 4)

}
