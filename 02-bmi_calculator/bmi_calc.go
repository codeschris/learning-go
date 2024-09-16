package main

import(
	"fmt"
)

func main(){
	var weight, height float64

	fmt.Print("Enter your weight (kg): ")
	fmt.Scanln(&weight)
	fmt.Print("Enter your height (m): ")
	fmt.Scanln(&height)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI is: %.2f \n", bmi)

	if bmi < 18.5 {
		fmt.Println("Category: Underweight")
	} else if bmi >= 18.5 && bmi <= 24.9 {
		fmt.Println("Category: Normal weight")
	} else if bmi >= 25 && bmi <= 29.9 {
		fmt.Println("Category: Overweight")
	} else {
		fmt.Println("Category: Obese")
	}
}