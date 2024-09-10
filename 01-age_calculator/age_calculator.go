package main

import (
	"fmt"
	"time"
)

func main(){
	var birthYear int	// One way to declare a variable
	fmt.Print("Enter your year of birth: ")
	fmt.Scanln(&birthYear)

	currentYear := time.Now().Year()	// Another way to declare variables
	age := currentYear - birthYear
	fmt.Printf("You are %d years old \n", age)
}