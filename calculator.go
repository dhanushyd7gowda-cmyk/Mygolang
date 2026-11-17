package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("AZIRO Calculator")

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Println("Result:", num1+num2)

	case "-":
		fmt.Println("Result:", num1-num2)

	case "*":
		fmt.Println("Result:", num1*num2)

	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Println("Result:", num1/num2)
		}

	default:
		fmt.Println("Invalid operator!")
	}
}
