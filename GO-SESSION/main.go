package main

import "fmt"

func main() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	fmt.Print(devide(10, 3))
}

func devide(num1, num2 int) int {
	if num2 == 0 {
		panic("Can't divide by zero")
	}
	result := num1 / num2
	return result
}
