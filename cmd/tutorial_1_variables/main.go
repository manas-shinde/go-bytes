package main

import "fmt"

func main(){
	fmt.Println("Namaste, Duniya!")

	// declare a integer variable (there are multiple format for integer declaration)
	var age int = 25
	fmt.Println("Age:", age)

	// there are multiple ways to declare a variable in Go
	favNumber := 11 // using short variable declaration
	fmt.Println("Favorite Number:", favNumber)

	// float32 and float64 are two types of floating point numbers in Go
	var floatNum1 float32 = 356151.14 // using explicit type declaration
	var floatNum2 = 2.718 // using type inference

	fmt.Println("Float Number 1:", floatNum1)
	fmt.Println("Float Number 2:", floatNum2)


	var fullName string = "Manas Shinde" // using explicit type declaration
	var fullName2 = `Mr.
	Manas Shinde 
	(using backticks)` // using backticks for multi-line string
	fmt.Println("Full Name:", fullName)
	fmt.Println("Full Name 2:", fullName2)

	fmt.Print("length of fullName: ", len(fullName), "\n")

	//single character in Go is represented as a rune (which is an alias for int32)
	var singleChar rune = 'M' // using single quotes for rune
	fmt.Println("Single Character:", singleChar)

	// boolean type
	var isActive bool = true // using explicit type declaration
	fmt.Println("Is Active:", isActive)	

	// constants in Go
	const pi float64 = 3.14159 // using const keyword for constant declaration
	fmt.Println("Value of Pi:", pi)

	// using multiple variable declaration
	var (
		city    string = "Pune";
		country string = "India"
	)
	// or using short variable declaration
	// city, country := "Pune", "India"
	fmt.Println("City:", city)
	fmt.Println("Country:", country)

}