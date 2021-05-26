// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

// Declaring variables at the package level uses declaration type 1 only.
var s int = 12

// If we want to declare multiple variables at the package level, then follow this method.
var (
	name  string  = "genzyy"
	age   int     = 100
	money float32 = 0.0
)

// Once a variable is declared, we have to use it. An undeclared variable will give an error.
// The length of the variable name determines the lifespan of the variable.

func main() {

	// Use this type when we know the variable but dont know what will
	// be its value.
	var i int
	i = 42
	i = 27
	fmt.Println(i)
	fmt.Printf("%v %T", i, i)
	// Variable type conversion
	h := float32(42)
	fmt.Println(h)
	fmt.Println()

	// Use this to have more control over the variable type.
	var j int = 42
	fmt.Println(j)
	fmt.Printf("%v %T", j, j)
	fmt.Println()

	// Let the compiler decide what will be its value.
	k := 36
	fmt.Println(k)
	fmt.Printf("%v %T", k, k)
	fmt.Println()
	l := 99.0
	fmt.Println(l)
	fmt.Printf("%v %T", l, l)
	fmt.Println()
	// By default the compiler will use float64
	// To get more control on it, use type 2 declaration.
	var m float32 = 9.0
	fmt.Println(m)
	fmt.Printf("%v %T", m, m)
	fmt.Println()

	// Shadowing
	// This means the local s variable shadows the value of the global s variable.
	fmt.Println(s)
	fmt.Println()
	var s int = 9
	fmt.Println(s)
	fmt.Println()
	// While variables are case sensitive, the case of the first letter of a variable has special meaning in Go.
	// If a variable starts with an uppercase letter, then that variable is accessible outside the package it
	// was declared in (or exported). If a variable starts with a lowercase letter, then it is only available
	// within the package it is declared in.
	var I int = 42
	fmt.Println(I)

	// To convert an integer or any number to string.
	var str string
	// This will give out the unicode for the number i=42 which is * (asterisk).
	str = string(i)
	// To convert the number into an actual string, use strconv.Itoa()
	str = strconv.Itoa(i)
	fmt.Println(str)
}
