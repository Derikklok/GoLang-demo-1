package main

import (
	"fmt"
)

func main() {
	// Go has three main data types
	/*
		bool
		numeric - int , float
		string
	*/

	var isStudent bool = true
	var age int = 23
	var height float32 = 195.56
	var name string = "Sachin"
	fmt.Println("Nice Work!")
	fmt.Println("Student or not\t:\t", isStudent)
	fmt.Println("Student Age\t:\t", age)
	fmt.Println(height)
	fmt.Println(name)

	// Arrays

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	arr2 := [5]int{10, 2, 56, 43, 4}
	fmt.Println(arr2)

	

}
