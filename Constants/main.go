package main

import (
	"fmt"
)

const (
	e = iota
	f
	g
)

func main() {
	const myConst int = 42
	fmt.Printf("%v %T\n", myConst, myConst)

	// The below line cannot be executed since it requires the calculation of the
	// Sin function which is only done during the runtime. And since we need to
	// assign a value to the constant during the compile time we don't have
	// any value to assign.
	//const myValue float64 = math.Sin(1.57)

	// We cant add a const and a variable of different types but we can do that kind of
	// thing the way given below.
	const a = 42
	var b int16 = 27
	fmt.Printf("%v %T\n", a+b, a+b)
	// The above line is interpreted as fmt.Printf("%v %T\n", 42 + b, 42 + b)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	// The constants in the package level gets to know what the
	// pattern is to declare the constants and therefore it gives them
	// the value accordingly.
}
