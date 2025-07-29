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
	var result, remainder, err = intDivision(num, den)
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

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, den int) (int, int, error) {
	var error error = nil
	if den == 0 {
		error = errors.New("Division by zero error")
		return 0, 0, error
	}
	var result int = num / den
	var remainder int = num % den

	return result, remainder, error
}
