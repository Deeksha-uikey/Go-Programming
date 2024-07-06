package main

import "fmt"

func main() {
	var a int
	var b int

	// Input from the user
	fmt.Println("Enter any number for a: ")
	fmt.Scan(&a)

	fmt.Println("Enter any number for b: ")
	fmt.Scan(&b)

	//Addition
	result := a + b
	println("\n result of a + b = %d\n", result)

	//substraction
	result1 := a - b
	println("\n result of a - b = %d\n", result1)

	//Multiplication
	result2 := a * b
	println("\n result of a * b = %d\n", result2)

	//Division
	result3 := a / b
	println("\n result of a / b = %d\n", result3)

	//Modulus
	result4 := a % b
	println("\n result of a %% b = %d\n", result4)

}
