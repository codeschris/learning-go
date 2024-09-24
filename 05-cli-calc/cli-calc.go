package main

import (
	"fmt"
)

func main(){
	var num1, num2 float64
	var operator string
	
	fmt.Println("CLI Calculator")
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil{
		fmt.Println("Invalid input")
		return
	}
	
	fmt.Print("Enter the operator (+, -, /, *): ")
	fmt.Scanln(&operator)
}