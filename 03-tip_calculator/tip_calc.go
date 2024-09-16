package main

import (
	"fmt"
)

func main(){
	var billAmount, tipPercentage float64

	fmt.Print("Enter the bill you have been charged: ")
	fmt.Scanln(&billAmount)
	fmt.Print("Enter the tip percentage: ")
	fmt.Scanln(&tipPercentage)

	tipAmount := (tipPercentage / 100) * billAmount
	totalBillAmount := billAmount + tipAmount

	fmt.Printf("Tip amount: Ksh %.2f\n", tipAmount)
	fmt.Printf("Total bill amount with tip: Ksh %.2f\n", totalBillAmount)
}