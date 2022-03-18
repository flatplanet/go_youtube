package main

// import other packages
import (
	"fmt"
)

func main(){

	// Variables and Constants
	// string, int, float32, bool

	//var firstName string = "John"

	// := shorthand variables

	//lastName := "Elder"
	// age := 41

	// Create but don't assign

	var fullName string
	var age int
	var price float32
	var tf bool

	fmt.Println(fullName, age, price, tf)


	fullName = "John Elder"
	age = 44
	price = 19.99
	tf = true

	fmt.Println(fullName, age, price, tf)

	// Multiple variable declaration

	var name1, name2 string = "Tim", "Mary"
	fmt.Println(name1, name2)

	name1 = "Fred"
	fmt.Println(name1, name2)

	
	// Constants
	const PIZZA string = "Pepperoni"
	fmt.Println(PIZZA)

	//PIZZA = "Cheese"
	//fmt.Println(PIZZA)

	// Multiple Constants

	const (
		PIZZA1 = "Peperoni"
		PIZZA2 = "Cheese"
		MYNUM = 77
	)

	fmt.Println(PIZZA2)


	// fmt.Println(firstName)

}
