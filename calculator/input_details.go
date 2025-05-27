package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputDetails() {

	// --- ✨ fmt.Print উদাহরণ ---
	fmt.Print("এই লাইনটি fmt.Print দিয়ে লেখা (new line নাই)।")
	fmt.Print(" এটা একই লাইনে যাবে।\n") // \n দিয়ে লাইন ব্রেক এনেছি

	// --- ✨ fmt.Println উদাহরণ ---
	fmt.Println("এই লাইনটি fmt.Println দিয়ে লেখা, শেষে অটোমেটিক new line দিবে।")
	fmt.Println("আরেকটা নতুন লাইন।")

	// --- ✨ fmt.Printf উদাহরণ ---
	name := "Santo"
	age := 25
	fmt.Printf("fmt.Printf দিয়ে: Name: %s, Age: %d\n", name, age)

	// --- ✨ fmt.Scan উদাহরণ ---
	var scanName string
	var scanAge int
	fmt.Print("fmt.Scan → Enter your name and age (space separated): ")
	fmt.Scan(&scanName, &scanAge)
	fmt.Println("Scan Output:", scanName, scanAge)

	// --- ✨ fmt.Scanln উদাহরণ ---
	var word1, word2 string
	fmt.Print("fmt.Scanln → Enter two words (press Enter after input): ")
	fmt.Scanln(&word1, &word2)
	fmt.Println("Scanln Output:", word1, word2)

	// --- ✨ fmt.Scanf উদাহরণ ---
	var age2 int
	var name2 string
	fmt.Print("fmt.Scanf → Enter age and name like this (30 : Rakib): ")
	fmt.Scanf("%d : %s", &age2, &name2)
	fmt.Printf("Scanf Output: Age: %d, Name: %s\n", age2, name2)
}

func input() {

	//fmt.Println("Md Santo")
	var name string
	var age int
	//fmt.Scan(&name)
	//fmt.Scanln(&name)
	fmt.Scanf("%s : %d", &name, &age)
	//fmt.Println("You name : ", name, " ", age)
	fmt.Printf("Your name : %s , Your age  %d\n ", name, age)

	// Input abar dewar jonno
	bufio.NewReader(os.Stdin).ReadString('\n')
	// Reader crate
	reader := bufio.NewReader(os.Stdin)
	sentance, _ := reader.ReadString('\n')
	println(sentance)

}
