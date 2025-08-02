package main

import "fmt"

/*
Pointers are one of the most powerful and fundamental features of the Go programming language.
They allow you to manipulate memory directly, create dynamic data structures, and improve program efficiency.
However, pointers can also be tricky to use, and incorrect usage can lead to memory leaks and other errors
*/
func main() {
	fmt.Println("This is a placeholder for the main function in tutorial 5 pointers.")
	// pointerExample1()
	// pointerExample2()
	// pointerExample3()
	// pointerExample4()
	pointerExample5()
}
func pointerExample1() {
	var x int = 10
	var p *int = &x                            // p is a pointer to x
	fmt.Println("Value of x:", x)              // Output: Value of x: 10
	fmt.Println("Address of x:", &x)           // Output: Address of x: 0x... // Address of x
	fmt.Println("Value of p:", p)              // Output: Value of p: 0x... // Address of x
	fmt.Println("Value pointed by p:", *p)     // Output: Value pointed by p: 10
	*p = 20                                    // Change the value of x through the pointer
	fmt.Println("New value of x:", x)          // Output: New value of x: 20
	fmt.Println("New value pointed by p:", *p) // Output: New value pointed by
}

func pointerExample2() {
	/*
		we declare a pointer variable ptr of type *int and use the new function to allocate memory for an integer value. We then print the memory address stored in ptr, which is the address of the newly allocated memory block. We also print the value of *ptr, which is the value stored at the memory address, which is initially set to 0.
		We then assign the value 10 to the memory location pointed to by ptr using the * operator.
	*/
	var p *int = new(int)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 10
	fmt.Println("Value pointed by p after assignment:", *p) // Output: Value pointed by p after assignment: 10
	fmt.Println("Address of p:", p)                         // Output: Address of p: 0x... // Address of p
}

func pointerExample3() {
	var x int = 5
	fmt.Println("Value of x before increment:", x) // Output: Value of x before increment: 5
	incrementByOne(&x)                             // Pass the address of x to the function incrementByOne
	fmt.Println("Value of x after increment:", x)  // Output: Value of x after increment: 6

}

func incrementByOne(x *int) {
	*x += 1 // Increment the value at the address pointed to by x
}

func pointerExample4() {
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	printArray(&arr) // Pass the address of the array to the function printArray
}

func printArray(arr *[5]int) {
	for i, v := range arr {
		fmt.Printf("Element %d: %d\n", i, v)
	}
}

func pointerExample5() {
	/*
		Functions are one of the most common use cases for pointers in Go. You can use pointers to pass arguments to functions by reference, which allows the function to modify the original parameter value. You can also use pointers to return values from functions by reference, which allows the function to modify a variable outside its scope.
	*/
	var x int = 5
	var y int = 10
	fmt.Println("Before swap: x =", x, ", y =", y)
	swap(&x, &y)                                  // Pass the addresses of x and y to the swap function
	fmt.Println("After swap: x =", x, ", y =", y) // Output: After swap: x = 10 , y = 5

	var z *int = doubleValue(&x)                  // Pass the address of x to the doubleValue function
	fmt.Println("Value of x after doubling:", *z) // Output: Value of x after doubling: 10
	fmt.Println("Address of x:", &x)              // Output: Address of x: 0x... // Address of x
	fmt.Println("Value pointed by z:", *z)        // Output: Value pointed by z: 10
	fmt.Println("Address of z:", z)               // Output: Address of z: 0x... // Address of x
}

func doubleValue(x *int) *int {
	/*
		This function takes a pointer to an integer as an argument and returns a pointer to an integer.
		It doubles the value of the integer pointed to by the input pointer and returns a pointer to the new value.
	*/
	result := *x * 2
	return &result
}
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
