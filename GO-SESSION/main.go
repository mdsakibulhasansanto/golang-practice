package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (float64, float64, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("First number: ")
	num1Str, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, "", fmt.Errorf("Failed to read first number: %v", err)
	}
	num1Str = strings.TrimSpace(num1Str)
	floatNum1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("Failed to convert first number: %v", err)
	}

	fmt.Print("Second number: ")
	num2Str, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, "", fmt.Errorf("Failed to read second number: %v", err)
	}
	num2Str = strings.TrimSpace(num2Str)
	floatNum2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("Failed to convert second number: %v", err)
	}

	fmt.Print("Operator (+, -, *, /): ")
	operator, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, "", fmt.Errorf("Failed to read operator: %v", err)
	}
	operator = strings.TrimSpace(operator)

	return floatNum1, floatNum2, operator, nil
}

func main() {
	// Panic handling using defer + recover
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	num1, num2, op, err := getInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("Cannot divide by zero")
		}
		result = num1 / num2
	default:
		panic("Invalid operator")
	}

	fmt.Printf("Result: %.2f\n", result)
}
