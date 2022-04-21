package main

// import other packages
import (
	"fmt"
	"strconv"
)


func main(){

	// String to Int Type Conversion
	var myNum string = "Bob"
	
	// Convert String to Integer
	myInt, err := strconv.Atoi(myNum)

	if err == nil {
		fmt.Println(myInt + 4)
	} else {
		fmt.Println("Hey! There was an error!", err)
	}
	
	


}
