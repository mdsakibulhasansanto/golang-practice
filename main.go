package main

import (
	"fmt"
	"net/http"
)

func main() {
	ch := make(chan string)

	link := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	for _, value := range link {
		go checkLink(value, ch)
	}

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	for link := range ch {
		go checkLink(link, ch)
	}
}

func checkLink(link string, ch chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println("This site down", link)
		ch <- link
		return
	}

	fmt.Println("This site up", link)
	ch <- link
}
