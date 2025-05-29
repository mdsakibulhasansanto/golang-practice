package main

import (
	"fmt"
	"strconv"
)

func typeConversion() {

	var rollStr, classStr string

	fmt.Print("Enter roll and class: ")
	n, scanErr := fmt.Scan(&rollStr, &classStr)
	if scanErr != nil {
		fmt.Println(" Scan Error:", scanErr)
		return
	}
	fmt.Println(" Input Count:", n)
	fmt.Println("Raw Input => Roll:", rollStr, ", Class:", classStr)

	//  String to int
	rollInt, err1 := strconv.Atoi(rollStr)
	if err1 != nil {
		fmt.Println(" Roll conversion error:", err1)
	} else {
		fmt.Println(" Roll (String to Int):", rollInt)
	}

	// String to float64
	classFloat, err2 := strconv.ParseFloat(classStr, 64)
	if err2 != nil {
		fmt.Println(" Class conversion error:", err2)
	} else {
		fmt.Println(" Class (String to Float64):", classFloat)
	}

	// String to bool (extra example)
	boolStr := "true"
	boolVal, err3 := strconv.ParseBool(boolStr)
	if err3 != nil {
		fmt.Println(" Bool conversion error:", err3)
	} else {
		fmt.Println(" Bool (String to Bool):", boolVal)
	}

	// Int to String
	rollToStr := strconv.Itoa(rollInt)
	fmt.Println(" Int to String:", rollToStr)

	// Float64 to String
	classToStr := fmt.Sprintf("%.2f", classFloat)
	fmt.Println(" Float64 to String:", classToStr)

	// Bool to String
	boolToStr := strconv.FormatBool(boolVal)
	fmt.Println(" Bool to String:", boolToStr)

	//  Int to Float64
	rollToFloat := float64(rollInt)
	fmt.Println(" Int to Float64:", rollToFloat)

	//  Float64 to Int
	floatToInt := int(classFloat)
	fmt.Println(" Float64 to Int:", floatToInt)

	//  Byte and Rune to String
	var myByte byte = 65
	var myRune rune = 66
	fmt.Println(" Byte to String:", string(myByte)) // Output: A
	fmt.Println(" Rune to String:", string(myRune)) // Output: B
}
