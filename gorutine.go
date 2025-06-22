package main

import "fmt"

/*
func main() {
	ch := make(chan string)
	fmt.Println("Md Sakibul Hasan Santo")
	go printMessage("world", ch)

	message := <-ch
	fmt.Println(message)
}

*/

func printMessage(message string, ch chan string) {
	fmt.Println("Hello", message)
	ch <- "Complete"
}
