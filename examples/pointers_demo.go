package main

import "fmt"

// Function that takes a pointer to modify the original value
func doubleValue(num *int) {
	*num = *num * 2 // Dereference pointer and modify the value
}

// Function that takes a regular parameter (copy)
func tryToDouble(num int) {
	num = num * 2 // This only modifies the copy, not original
	// Note: this change doesn't affect the original variable
	fmt.Printf("Inside function: num = %d (but original won't change)\n", num)
}

// Function that swaps two values using pointers
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	fmt.Println("=== Pointer Basics ===")

	// 1. Basic pointer declaration and usage
	x := 42
	ptr := &x // Declare and assign pointer in one line

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Pointer ptr points to: %p\n", ptr)
	fmt.Printf("Value at pointer: %d\n", *ptr)

	// 2. Modifying through pointer
	fmt.Println("\n=== Modifying Through Pointer ===")
	*ptr = 100 // Change x through the pointer
	fmt.Printf("After *ptr = 100, x is now: %d\n", x)

	// 3. Function with pointer parameter
	fmt.Println("\n=== Function with Pointer ===")
	y := 5
	fmt.Printf("Before doubleValue: y = %d\n", y)
	doubleValue(&y) // Pass address of y
	fmt.Printf("After doubleValue: y = %d\n", y)

	// 4. Function without pointer (pass by value)
	fmt.Println("\n=== Function without Pointer ===")
	z := 5
	fmt.Printf("Before tryToDouble: z = %d\n", z)
	tryToDouble(z)                               // Pass copy of z
	fmt.Printf("After tryToDouble: z = %d\n", z) // z unchanged!

	// 5. Swapping values with pointers
	fmt.Println("\n=== Swapping Values ===")
	a, b := 10, 20
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)

	// 6. Nil pointers - more realistic example
	fmt.Println("\n=== Nil Pointers ===")

	// Function that might return nil
	findNumber := func(search int) *int {
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			if num == search {
				return &num
			}
		}
		return nil // Not found
	}

	// Safe pointer usage with nil check
	result := findNumber(10) // Looking for 10 (doesn't exist)
	if result == nil {
		fmt.Println("Number not found - pointer is nil")
	} else {
		fmt.Printf("Found number: %d\n", *result)
	}

	result2 := findNumber(3) // Looking for 3 (exists)
	if result2 != nil {
		fmt.Printf("Found number: %d\n", *result2)
	}

	// 7. Pointer to string
	fmt.Println("\n=== String Pointers ===")
	message := "Hello"
	strPtr := &message
	fmt.Printf("Original: %s\n", message)
	*strPtr = "Hello, Go!"
	fmt.Printf("Modified through pointer: %s\n", message)
}
