package main

import (
	"errors"  // errors module
	"fmt"     // formated output and inputs
	"strconv" // string operations and convertions
)

func main() {
	fmt.Println("Hello , 世界") // print Hello World

	//
	// variable declaration
	fmt.Println("\n**** Variable Declarations ****")
	var a int
	var b int = 10
	a = 15
	// OR
	c := 20
	fmt.Println("Value a: ", a)
	fmt.Println("Value b: ", b)
	fmt.Println("Sum: ", sum(a, b))
	fmt.Println("Value c: ", c)

	//
	// array declarations and initialization
	fmt.Println("\n**** Array Declarations ****")
	arr1 := []int{1, 2, 3, 4, 5}
	var arr2 [5]int // dynamic sized arrays
	arr2[1] = 1
	fmt.Println("arr1: ", arr1)
	arr1 = append(arr1, 7) //	Appending a new element to an array
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	for index, value := range arr1 { // Using range to loop through the array items
		fmt.Println("index:", index, "value:", value)
	}

	// map declarations
	fmt.Println("\n**** map declaration ****")
	dummyMap := make(map[string]string)
	dummyMap["a"] = "test 1"
	dummyMap["b"] = "test 2"
	fmt.Println(dummyMap)
	fmt.Println("Value of a:", dummyMap["a"])
	for key, value := range dummyMap {
		fmt.Println("key:", key, "value:", value)
	}

	//
	// Method declarations
	fmt.Println("\n**** Method Declarations ****")
	fmt.Println("Method to check if a value is greater or not")
	result, error := isGreater(6, 11) // equivalent for spread operation
	fmt.Println("Result: ", result)
	fmt.Println("Error: ", error)

	//
	// Struct declaration
	fmt.Println("\n**** Struct Declarations ****")
	type person struct {
		name string
		age  int
	}
	p1 := person{name: "John", age: 23}
	fmt.Println("Struct Person: ", p1)
	fmt.Println("Age of p1: ", p1.age)

	//
	// Pointers
	fmt.Println("\n**** Pointers ****")
	test := 5
	fmt.Println("Variable test: ", test)
	fmt.Println("Pointer of test: ", &test)
	increment(&test)
	fmt.Println("Variable test: ", test)
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

func increment(i *int) {
	*i++
}
