package main

// import other packages
import (
	"fmt"
)

func main(){

	// Logic
	var firstName string = "Mary"

	if firstName == "John" {
		fmt.Println("Hello", firstName)
	} else if firstName == "Tim" {
		fmt.Println("What's Up", firstName, "!")
	} else {
		fmt.Println("Sorry", firstName, "I Don't Know You!")
	}

	myNum := 12
	if myNum > 30 {
		fmt.Println(myNum, "is greater than 30")
	} else {
		fmt.Println(myNum, "is less than 30")
	}
}
