package main

import (
	"fmt"
)

func main() {
	var num1, num2 float32

	// Addition
	fmt.Println("\t--- Addition ---")
	fmt.Print("\tEnter two numbers for addition: ")
	fmt.Scanln(&num1, &num2)
	resultSum := num1 + num2
	fmt.Printf("\tThe sum of %.2f and %.2f == %.2f\n\n", num1, num2, resultSum)

	// Subtraction
	fmt.Println("\t--- Subtraction ---")
	fmt.Print("\tEnter two numbers for subtraction: ")
	fmt.Scanln(&num1, &num2)
	resultSub := num1 - num2
	fmt.Printf("\tThe subtraction of %.2f and %.2f == %.2f\n\n", num1, num2, resultSub)
	// Multiplication
	fmt.Println("\t--- Multiplication ---")
	fmt.Print("\tEnter two numbers for multiplication: ")
	fmt.Scanln(&num1, &num2)
	resultMultiplication := num1 * num2
	fmt.Printf("\tThe multiplication of %.2f and %.2f == %.2f\n\n", num1, num2, resultMultiplication)

	// Division
	fmt.Println("\t--- Division ---")
	fmt.Print("\tEnter two numbers for division: ")
	fmt.Scanln(&num1, &num2)
	if num2 != 0 {
		resultDivision := num1 / num2
		fmt.Printf("\tThe division of %.2f by %.2f == %.2f\n", num1, num2, resultDivision)
	} else {
		fmt.Println("\tError: Cannot divide by zero.")
	}

	fmt.Println("\t--- End calculate ---")
}
