package main
import (
	"fmt"
	"strings"
)

func main(){
	// String Manipulation
	var name string = " John elder "
	fmt.Println(name)

	// Uppercase
	fmt.Println("Upper", strings.ToUpper(name))

	// Lowercase
	fmt.Println("Lower", strings.ToLower(name))

	// Titlecase
	fmt.Println("Title:", strings.Title(name))


	// Trim
	fmt.Print("Trim:", name, "Something")
	fmt.Println(" ")
	fmt.Print("Trim:", strings.TrimRight(name, " "), "Something")
	fmt.Println(" ")
	fmt.Print("Trim:", strings.TrimLeft(name, " "), "Something")
}

