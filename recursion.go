package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// Cual es mas rapido?
	// Realizar pruebas de tiempo

	start := time.Now()
	fmt.Println("Factorial of 25 is:", recursivefactorial(25))
	elapsed := time.Since(start)
	fmt.Println("Tiempo tomado por recursivefactorial:", elapsed)
	start = time.Now()
	fmt.Println("Factorial of 25 is:", forFactorial(25))
	elapsed = time.Since(start)
	fmt.Println("Tiempo tomado por forFactorial:", elapsed)
	fmt.Println("Factorial of 25 is:", forFactorial(25))



	fmt.Println("El mas rapido es forFactorial")
}

func recursivefactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursivefactorial(n-1)
}

func forFactorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
