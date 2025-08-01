package main

import (
	"fmt"
	"go-bytes/cmd/tutorial_4_struct_interface/areacalc"
)

type Employee struct {
	name       string
	age        int
	employeeId int
	city       string
}

func main() {
	emp1 := Employee{
		name:       "Manas Shinde",
		age:        25,
		employeeId: 12345,
		city:       "Pune",
	}

	var emp2 Employee // zero value of Employee struct

	emp1.printEmployeeDetails()
	printEmployee(&emp2) // Passing emp2 by reference to modify it
	fmt.Println("After modifying emp2:")
	emp2.printEmployeeDetails()

	fmt.Println("Calculating areas and perimeters of shapes:")
	// Call the area and perimeter calculations
	areacalc.Main()

}
func printEmployee(e *Employee) {
	fmt.Println(*e)
	e.name = "John Doe"
}

/*
A method augments a function by adding an extra parameter section immediately after the func keyword that accepts a single argument.
This argument is called a receiver.
*/
func (e Employee) printEmployeeDetails() {
	fmt.Printf("Name: %s\n", e.name)
	fmt.Printf("Age: %d\n", e.age)
	fmt.Printf("Employee ID: %d\n", e.employeeId)
	fmt.Printf("City: %s\n", e.city)
}
