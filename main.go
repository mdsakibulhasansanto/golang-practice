package main

import "fmt"

func main() {
	var name_list [2]int
	name_list[0] = 10
	//name_list[1] = 20

	num_list := [4]int{10, 20, 30, 40}

	array_un := [...]int{10, 30, 40, 50, 30}

	fmt.Println(num_list)
	fmt.Println(name_list)
	fmt.Println(array_un)

}
