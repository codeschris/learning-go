package main

import (
	"fmt"
)

func celsiusToFahrenheit(celcius float64) float64 {
	return (celcius * 9/5) + 32
}

func fahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5/9
}

func main(){
	fmt.Println("Temperature Converter")
	fmt.Println("1. Celcius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celcius")
	fmt.Print("Enter the conversion type (1 or 2): ")
	
	var choice int
	_, err := fmt.Scanln(&choice) // Research on this for clarity
	if err != nil{
		fmt.Println("Invalid input")
		return
	}
	if choice != 1 && choice != 2{
		fmt.Println("Choose either option 1 or 2")
	}
	
	var temp float64
	fmt.Print("Enter the temperature value: ")
	_, err = fmt.Scanln(&temp)
	if err != nil{
		fmt.Println("Invalid temperature input")
		return
	}
	
	switch choice{
		case 1:
			result := celsiusToFahrenheit(temp)
			fmt.Printf("Temperature (%.2f Celsius) is %.2f Fahrenheit\n", temp, result)
		case 2:
			result := fahrenheitToCelcius(temp)
			fmt.Printf("Temperature (%.2f F) is %.2f Celsius degrees\n", temp, result)
	}
}