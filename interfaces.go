package main

import "fmt"

type engine struct {
	name    string
	gallons uint8
	owner   company
}

type company struct {
	name    string
	revenue uint8
}

type furniture interface {
	getPrice() uint8
}

type chair struct {
	price uint8
}

func (e chair) getPrice() uint8 {
	return e.price * 10
}

type table struct {
	price uint8
}

func (e table) getPrice() uint8 {
	return e.price * 2
}

func (a engine) greeting() string {
	return a.owner.name + " " + a.name + " Welcomes you!"
}

func bill(a furniture) {
	fmt.Print(a.getPrice(), "\n")
}

func main() {
	myString := "Hello guys"
	for i, v := range myString {
		fmt.Printf("%d %c\n", i, v)
	}
	var myEngine engine = engine{"Toyota", 200, company{"Mitsubiushi", 200}}
	fmt.Print(myEngine.gallons, " ", myEngine.owner.name, "\n")
	greeting := myEngine.greeting()
	fmt.Print(greeting + "\n")
	var chairrr chair = chair{10}
	var tablee table = table{20}
	bill(chairrr)
	bill(tablee)
}
