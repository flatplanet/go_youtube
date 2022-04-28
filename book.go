package main
import (
"fmt"
"strings"
)

func main(){
	var name string = "John "
	fmt.Print(strings.TrimRight(name, " "), "Stuff")
	fmt.Println("")
	name = "John Elder"
	// Lower Case
	fmt.Println(strings.ToLower(name))

	// Upper Case
	fmt.Println(strings.ToUpper(name))

	// Title
	fmt.Println(strings.Title(name))
}

