package main

import "fmt"

func main() {
	fmt.Println()
	grades := [3]int{93, 82, 71} // Contiguous in memory -> faster compilation.
	// Have enough memory to store the data that has been passed to the array.
	newgrades := [...]int{93, 82, 71, 60, 59, 48, 37}
	fmt.Println(newgrades)
	fmt.Println(grades)
	fmt.Println(len(grades))
	fmt.Println(len(newgrades))
	// Another way of declaring an array and a 2D array too.
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
	a := [3]int{1, 2, 3}
	b := a
	b[1] = 5
	// Copying an array into another variable creates a copy of that array and doesn't link it to the same
	// memory like some programming languages do.
	fmt.Println(a)
	fmt.Println(b)
	// Array type.
	fmt.Printf("%v %T\n", a, a)
	// Storing the array memory location to another variable.
	d := &a
	fmt.Println(d)
	// Type of the addressed array variable.
	fmt.Printf("%v %T\n", d, d)
	// Working with slices.

	m := []int{1, 2, 3}
	p := m
	// Variable p will be referencing to the array m as this is an exception in slices.
	fmt.Println(m)
	fmt.Printf("%v %T\n", m, m)
	fmt.Println(p)
	fmt.Println(len(m))
	fmt.Println(len(p))

}
