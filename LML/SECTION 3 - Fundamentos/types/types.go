package main

import (
	"fmt"
	"math"
	"reflect"
	
)

func main() {
	
	/* Whole numbers */
	fmt.Println(1, 2, 10000)
	fmt.Println("The literal type is", reflect.TypeOf(320000))
	
	/* No sign (only positives)... uint8 uint16 uint32 uint64*/
	var b byte = 255
	fmt.Println("The byte is", reflect.TypeOf(b))
	
	
	/* With signal... int8 int16 int32 int64 */
	i1 := math.MaxInt64
	fmt.Println("The maximum value of int is", i1)
	fmt.Println("i1 type is", reflect.TypeOf(i1))
	
}