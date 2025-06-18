package main

import "fmt"

func array_details() {
	var name_list [2]int
	num_list := [4]int{10, 20, 30, 40}
	array_un := [...]int{10, 20, 30, 40, 50, 30}

	var sum int
	for i := 0; i < len(array_un); i++ {
		sum += array_un[i]
	}

	fmt.Println("num_list   :", num_list)
	fmt.Println("name_list  :", name_list)
	fmt.Println("array_un   :", array_un, "Length:", len(array_un))
	fmt.Println("Sum        :", sum)
}
