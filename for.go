package main

// import other packages
import (
	"fmt"
)

func main(){

	// For Loops
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}
	
	// looping an array with the Range Keyword
	var names = [3] string {"John", "April", "Wes"}
	for _, name := range names {
		fmt.Println(name)
	}

}
