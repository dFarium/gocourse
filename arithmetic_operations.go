package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Modulus (remainder): ", result)

	const pi float64 = 22.0 / 7.0
	fmt.Println("Approximation of Pi (floating-point division): ", pi)

	//exponentiation
	// result = int(math.Pow(float64(a), float64(b)))
	var value float64 = math.Pow(100.0, 2000.0)
	fmt.Println("Exponentiation (a^b): ", value)
}
