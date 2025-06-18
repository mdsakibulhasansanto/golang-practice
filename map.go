package main

import "fmt"

func map_details() {

	/**

	// Map crate
	name_list := make(map[int]string)

	map1 := map[int]string{
		1: "roll 1",
		2: "roll 2",
	}

	// Value add
	name_list[1] = "Santo"
	name_list[2] = "Sakib"
	name_list[3] = "Nurul"

	// Value show
	fmt.Println("Key 1:", name_list[1])

	// map  uses loop
	fmt.Println("Loop over name_list:")
	for index, value := range name_list {
		fmt.Println(index, "=>", value)
	}

	// map1
	fmt.Println("Loop over map1:")
	for key, val := range map1 {
		fmt.Println(key, "=>", val)
	}

	for i := 0; i <= len(map1); i++ {

		fmt.Println(map1[i])
	}

	*/

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	indexRemove := map[int]bool{

		2: true,
		4: true,
		6: true,
		8: true,
	}

	newSlice := make([]int, 0, len(slice))

	for index, value := range slice {

		if !indexRemove[index] {

			newSlice = append(newSlice, value)
		}
	}

	fmt.Println(newSlice)
}
