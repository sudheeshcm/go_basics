package main

import (
	"errors"  // errors module
	"fmt"     // formated output and inputs
	"strconv" // string operations and convertions
)

func main() {
	fmt.Println("Hello , 世界") // print Hello World

	// variable declaration
	var a int
	var b int = 10
	a = 15

	// array declarations and initialization
	arr1 := [5]int{1, 2, 3, 4, 5}
	var arr2 [5]int
	arr2[1] = 1

	fmt.Println("\narr1: ", arr1)
	fmt.Println("arr1: ", arr2)
	fmt.Println("Sum: ", sum(a, b))
	fmt.Println("\n**** Method to check if a value is greater or not ****")
	result, error := isGreater(6, 11) // equivalent for spread operation
	fmt.Println("Result: ", result)
	fmt.Println("Error: ", error)
}

func sum(a int, b int) int {
	return a + b
}

func isGreater(a int, b int) (bool, error) {
	// Control flows
	if a > b {
		return true, nil
	} else if a == b {
		return false, errors.New("Values are equal")
	} else {
		fmt.Println("Value: ", a)
		fmt.Println("Value compared with: ", b)
		return false, errors.New("Value is smaller than " + strconv.Itoa(b))
	}
}
