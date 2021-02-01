package main

import "fmt"

func main() {
	
	fmt.Print(SquareSum([]int{1, 2}))
	
}

func SquareSum(numbers []int) int {
	
	sum := 0
	
	for _, number := range numbers {
		
		sum += number * number
	}
	
	return sum
}