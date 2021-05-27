package main

import "fmt"

// everytime we shift a number we multiply/divide the number based on the direction of the operator
// that we have used in the expression.

const (
	_ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
}
