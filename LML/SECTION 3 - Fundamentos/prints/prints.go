package main

import "fmt"

func main() {
	
	fmt.Print("Same")
	fmt.Print(" Line.")
	fmt.Println(" Same line still.")
	fmt.Println("New line.")
	fmt.Println("New line.")
	
	x := 3.14159
	
	/* Concatenating and casting */
	fmt.Println("The value of x is " + fmt.Sprint(x))
	fmt.Println("The value of x is", x)
	fmt.Printf("The value of x is %.2f\n", x)
	
	a := 1
	b := 1.9999
	c := false
	d := "Hey!"
	
	fmt.Printf("%d\n%f\n%t\n%s", a, b, c, d)
	
}
