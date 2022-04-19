package main

// import other packages
import (
	"fmt"
)

type Person struct {
	name string
	age int
	favColor string
	weight int
}

func main(){

	// Structs
	var person1 Person

	//person1
	person1.name = "John"
	person1.age = 44
	person1.favColor = "Blue"
	person1. weight = 185

	fmt.Println("Person 1: ", person1)
	fmt.Println("Person 1 Name:", person1.name)

	// change and item
	person1.name = "John Elder"
	fmt.Println("Person 1 Name:", person1.name)	

	// shorthand person2
	person2 := Person {
		name: "Sally",
		age: 29,
		favColor: "Pink",
		weight: 115,
	}

	fmt.Println("Person 2: ", person2)
	fmt.Println("Person 2 Name: ", person2.name)

}
