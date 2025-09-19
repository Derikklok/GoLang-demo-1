package main

// Means that the file beongs to the main package
import (
	"fmt"
)

// fmt - formatted input output

var speed float32 = 234.45
var name string = "sachin"

func main() {

	var age int

	fmt.Println("Hello World!")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(speed)
}
