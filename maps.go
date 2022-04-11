package main

// import other packages
import (
	"fmt"
)


func main(){

	// Maps var pizza = map [key_type]value_type{...}
	var pizza = map[string]string{
		"John":"Pepperoni",
		"Mary":"Cheese",
		"Tim":"Mushroom"}

	fmt.Println(pizza)
	fmt.Println(pizza["John"])

	// Shorthand
	// toppings := map[string]string{"John":"Pepperoni"}

	// Change an Item
	pizza["John"] = "Peppers"
	fmt.Println(pizza["John"])

	// Delete Items
	delete(pizza, "Tim")
	fmt.Println(pizza)
	
	// Iterate Over Maps
	for key, value := range pizza {
		fmt.Println(key, value)
	}	


}
