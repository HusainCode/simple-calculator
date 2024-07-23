package main

import "fmt"

// Addition
func add(a, b float64) float64 {
	return a + b
}

// Subtraction
func subtract(a, b float64) float64 {
	return a - b
}

// Multiplication
func multiply(a, b float64) float64 {
	return a * b
}

// Division
func divide(a, b float64) float64 {
	if b != 0 {
		return a / b
	}
	fmt.Println("Error: Division by zero.")
	return 0
}

// Perform calculation
func calculate(num1, num2 float64, operator string) {
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, add(num1, num2))
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, subtract(num1, num2))
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, multiply(num1, num2))
	case "/":
		result := divide(num1, num2)
		if num2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
		}
	default:
		fmt.Println("Error: Invalid operator.")
	}
}

func main() {
	// Variables to hold the calculation values
	var num1, num2 float64
	// Variable to hold the operator sign
	var operator string

	// Prompt the user to enter the first number
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	// Prompt the user to enter the operator sign
	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Prompt the user to enter the second number
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	// Perform the calculation
	calculate(num1, num2, operator)
}
