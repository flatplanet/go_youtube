package main

// import other packages
import (
	"fmt"
)

func addIt(num1, num2 int){
	return (num1 + num2)
}

func myFunction(firstName string, age int){
	fmt.Println("This is my Function!")
	fmt.Println("Hello", firstName, "You are", age, "Years Old!")
}

func main(){

	// Functions
	myFunction("John", 44)

	fmt.Println(addIt(2,3))



}
