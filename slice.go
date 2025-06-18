package main

import (
	"fmt"
	"slices" // Available in Go 1.21+
)

func slice_details() {

	// 1. Create a slice using literal
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice1)

	// 2. Append values to the slice
	slice1 = append(slice1, 7, 8)
	fmt.Println("After append:", slice1)

	// 3. Append another slice
	more := []int{9, 10}
	slice1 = append(slice1, more...)
	fmt.Println("After appending another slice:", slice1)

	// 4. Check if a value exists using slices.Contains (Go 1.21+)
	fmt.Println("Contains 10?", slices.Contains(slice1, 10))

	// 5. Create slice using make() with length and capacity
	made := make([]int, 3, 5)
	made[0] = 10 // [0, 0, 0]
	fmt.Println("Slice using make():", made, "Len:", len(made), "Cap:", cap(made))

	// 6. Slice slicing (sub-slicing)
	sub := slice1[2:5] // index 2 to 4
	fmt.Println("Sub-slice [2:5]:", sub)

	// 7. Loop through a slice
	fmt.Println("Loop through slice:")
	for index, value := range slice1 {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// 8. Remove element by index (e.g., index 3)
	indexToRemove := 3
	slice1 = append(slice1[:indexToRemove], slice1[indexToRemove+1:]...)
	fmt.Println("After removing index 3:", slice1)

	// 9. Copy a slice
	copied := make([]int, len(slice1))
	copy(copied, slice1)
	fmt.Println("Copied slice:", copied)

	// 10. Reverse the slice using slices.Reverse (Go 1.21+)
	slices.Reverse(slice1)
	fmt.Println("Reversed slice:", slice1)

	// 11. Sort a slice using slices.Sort (Go 1.21+)
	slices.Sort(slice1)
	fmt.Println("Sorted slice:", slice1)

	// 12. Compare two slices using slices.Equal
	fmt.Println("Is equal to copied slice?", slices.Equal(slice1, copied))

	// 13. Multi-dimensional slice (2D slice)
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("2D Slice:", matrix)

	// 14. Slice is a reference type
	ref := slice1
	ref[0] = 100
	fmt.Println("Modified ref slice:", ref)
	fmt.Println("Original slice also changed:", slice1)
}
