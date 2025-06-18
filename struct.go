package main

import (
	"fmt"
)

func struhct_details() {

	type Adress struct {
		city   string
		state  string
		number int
	}

	type Person struct {
		name   string
		age    int
		adress Adress
	}

	p1 := Person{
		name: "Santo",
		age:  21,
		adress: Adress{
			city:   "Bangladesh",
			state:  "Dhaka",
			number: 12345,
		},
	}

	fmt.Println(p1)

}
