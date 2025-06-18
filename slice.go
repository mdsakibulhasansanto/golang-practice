package main

import (
	"fmt"
	"slices"
)

func slice_details() {

	slice1 := []int{1, 2, 3, 4, 5}

	slice1 = append(slice1, 7, 8)

	fmt.Println(slice1)
	fmt.Println(slices.Contains(slice1, 10))
}
