package main

import "fmt"

func CalculateGrade() {
	var marks int
	fmt.Print("Enter your marks: ")
	fmt.Scanln(&marks)

	switch {
	case marks < 40:
		fmt.Println("F grade")
	case marks < 50:
		fmt.Println("D grade")
	case marks < 60:
		fmt.Println("C grade")
	case marks < 70:
		fmt.Println("B grade")
	case marks < 80:
		fmt.Println("A grade")
	default:
		fmt.Println("A+ grade")
	}
}
