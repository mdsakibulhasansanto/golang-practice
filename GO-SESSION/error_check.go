package main

import (
	"fmt"
)

func errorCheck() {

	var num1, num2 int

	fmt.Print("Enter two numbers: ")
	n, err := fmt.Scan(&num1, &num2)

	if err != nil {
		fmt.Println(" Error reading input:", err)
		return
	}

	fmt.Printf("You entered (%d values): %d and %d\n", n, num1, num2)
}
