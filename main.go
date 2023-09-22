package main

import (
	"fmt"
	"math"
)

func Mod(num1 float64, num2 float64) float64 {
	return math.Mod(num1, num2)
}

func Power(num1 float64, num2 float64) float64 {
	return math.Pow(num1, num2)
}

func division(num1 float64, num2 float64) float64 {
	return num1 / num2
}

func subtract(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1 float64, num2 float64) float64 {
	return num1 / num2
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter an operator (+, -, *, /, %, ^): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	result := 0.0

	switch operator {
	case "+":
		result = addition() // DHIRAJ TODO
	case "-":
		result = subtract(num1, num2) // Randeep TODO
	case "*":
		result = multiply(num1, num2) // Abhijeet TODO
	case "/":
		if num2 != 0 {
			result = divison(num1, num2) // Mohammad Zubiyan Shaikh
		} else {
			fmt.Println("Error: Division by zero.")
			return
		}
	case "%": 
		if num2 != 0 {
			result = Mod(num1, num2)
		} else {
			fmt.Println("Error: Division by zero.")
			return
		}
	case "^":
		result = Power(num1, num2) // Abujar Shaikh
	default:
		fmt.Println("Error: Invalid operator")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
