package reciverfunc

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Reciver() {

	fmt.Println(p.Name)

}

func (p *Person) Update() {

	p.Name = "Nuha"
	fmt.Println(p.Name)
}
