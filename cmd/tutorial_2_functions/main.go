package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "String to print"
	printMe(printValue)

	var num int = 10
	var den int = 3
	var result, remainder, err = intDivision(num, den, printMe)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result of the integer division of %d / %d = %d\n", num, den, result)
	} else {
		fmt.Printf("Result of %d / %d = %d and remainder is %d\n", num, den, result, remainder)
	}

	// var res, rem, error = intDivision(num, 0)
	// switch {
	// case error != nil:
	// 	fmt.Println("Error:", error.Error())
	// case rem == 0:
	// 	fmt.Printf("Result of the integer division of %d / 0 = %d\n", num, res)
	// default:
	// 	fmt.Printf("Result of %d / 0 = %d and remainder is %d\n", num, res, rem)
	// }

	// Closure example
	fmt.Println("Closure example:")
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// Fibonacci example
	fmt.Println("Fibonacci example (using closure):")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, den int, printMe func(string)) (int, int, error) {
	var error error = nil
	if den == 0 {
		error = errors.New("Division by zero error")
		return 0, 0, error
	}
	var result int = num / den
	var remainder int = num % den

	// Functions are values too. They can be passed around just like other values.
	printMe("Inside intDivision function")

	return result, remainder, error
}

// Closure in Go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		ret := a
		a, b = b, a+b
		return ret
	}
}
