package main

import "fmt"

func main() {

	var num int = 10 // Variable declaration with type and value

	// Multiple variable declaration (only one used here)
	var (
		name string = "Santo"
	)

	names := "santo" // Variable declaration without type (type inferred as string)

	var p *int = &num // Pointer declaration, storing the address of 'num'

	const (
		low = iota + 1 // iota = 0
		medium
		high
	)

	// Printing the values
	fmt.Println(low, medium, high)
	fmt.Println("Value of num:", num)             // Prints the value 10
	fmt.Println("Value of name:", name)           // Prints "Santo"
	fmt.Println("Address of num (pointer p):", p) // Prints memory address like 0xc000012078
	fmt.Println("Value of names:", names)         // Prints "santo"

	// Pointer dereference to get the value from address
	// Dereferencing মানে হলো pointer যেই মেমোরি address ধরে রেখেছে, সেখানে গিয়ে ভ্যালু পড়া
	fmt.Println("Value at address stored in p:", *p) // Dereferencing pointer p → gives 10
}
