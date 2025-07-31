package main

import (
	"fmt"
)

func main() {
	// arrays()
	slices()
}

func arrays() {
	/*
		An array's length is part of its type, so arrays cannot be resized.
		This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
	*/
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Printf("Memory address of first element: %p\n", &primes[0])
}
func slices() {
	/*
		An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
		Slices are just wrappers around arrays, providing a more convenient way to work with sequences of elements.
		Unlike arrays, slices can grow and shrink as needed.
		In practice, slices are much more common than arrays.
	*/
	var s []int
	s = append(s, 1)
	fmt.Printf("Length of slice: %d\n", len(s))
	fmt.Printf("Capacity of slice: %d\n", cap(s))
	s = append(s, 2)
	s = append(s, 3)
	fmt.Printf("After appending elements:\n")
	fmt.Printf("Length of slice: %d\n", len(s))
	fmt.Printf("Capacity of slice: %d\n", cap(s))
	fmt.Println(s)
	fmt.Printf("Memory address of first element: %p\n", &s[0])

	var x []int
	fmt.Println(x, len(x), cap(x))

	//default value of a slice is nil
	if x == nil {
		fmt.Printf("Slice is nil\n")
	}

	//creating a slice using make
	a := make([]int, 5, 10)
	fmt.Printf("%s len=%d cap=%d %v\n", "a", len(a), cap(a), a)

	b := a[0:2]
	fmt.Printf("%s len=%d cap=%d %v\n", "b", len(b), cap(b), b)

	//append the data into the slice
	b = append(b, 3, 4, 5)
	fmt.Printf("%s len=%d cap=%d %v\n", "b", len(b), cap(b), b)

}

func range_in_go(s []uint8) {
	/*
		The range keyword in Go is used to iterate over elements in various data structures like arrays, slices, maps, and strings.
		It provides a convenient way to access both the index and value of each element.
	*/
	for i, v := range s {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

}
