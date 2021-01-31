package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	
	/* float casting */
	fmt.Println(x / float64(y))
	
	grade := 6.9
	final_grade := int(grade)
	fmt.Println(final_grade)
	
	/* ASCII casting */
	fmt.Println("test " + string(97))
	
	
	/* int -> string */
	fmt.Println("test " + strconv.Itoa(123))
	
	/* string -> int */
	num, _ := strconv.Atoi("230")
	fmt.Println(num - 100)
	
	/* str -> boolean */
	boolean, _ := strconv.ParseBool("false")
	if boolean{
		
		fmt.Print(boolean)
	
	} else {
		
		fmt.Print("Boolean parsing failed!")
	}
	
}
