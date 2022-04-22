package main

// import other packages
import (
	"fmt"
	"strconv"
)


func main(){

	// Int to String Type Conversion
	var myNum int = 41
	myNum2 := strconv.Itoa(myNum)
	fmt.Println(myNum2 + "Bob")

	// Second way
	var myNum3 int = 100
	myNum4 := fmt.Sprintf("%v", myNum3)
	fmt.Println(myNum4 + "Bob")


}
