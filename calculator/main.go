package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//fmt.Println("Md Santo")
	var name string
	var age int
	//fmt.Scan(&name)
	//fmt.Scanln(&name)
	fmt.Scanf("%s : %d", &name, &age)
	//fmt.Println("You name : ", name, " ", age)
	fmt.Printf("Your name : %s , Your age  %d\n ", name, age)

	bufio.NewReader(os.Stdin).ReadString('\n')
	// Reader crate
	reader := bufio.NewReader(os.Stdin)
	sentance, _ := reader.ReadString('\n')
	println(sentance)

}
