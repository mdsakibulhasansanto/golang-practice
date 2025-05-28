package main

import "fmt"

func sumNumber() {
	sum := 0

	// 0-10  sum
	fmt.Println(" Adding numbers from 0 to 10:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Adding: %d\n", i)
		sum += i
	}
	fmt.Println("Total Sum:", sum)

	//  Odd Numbers
	fmt.Println("\n Odd numbers from 0 to 10:")
	for x := 0; x <= 10; x++ {
		if x%2 != 0 {
			fmt.Println("Odd number:", x)
		}
	}

	//  Even Numbers - while loop
	fmt.Println("\n Even numbers from 0 to 10:")
	z := 0
	for z <= 10 {
		if z%2 == 0 {
			fmt.Println("Even number:", z)
		}
		z += 2
	}

	// infinite loop
	s := 0
	for {
		println("Number : ", s)
		s++
		if s > 5 {
			break
		}
	}
}
