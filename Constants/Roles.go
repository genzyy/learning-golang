package main

import "fmt"

// The value of the constants should be calculable during the compile time only.
const (
	isAdmin = 1 << iota
	isHeadQuarters
	canSeeFinance

	canSeeAfrica
	canSeeAsia
	canSeeNA
	canSeeSA
)

func main() {
	var roles byte = isAdmin | canSeeFinance | canSeeAsia
	fmt.Println(roles)
	fmt.Println(isAdmin)
	fmt.Println(isHeadQuarters)
	fmt.Println(canSeeFinance)
	fmt.Println(canSeeAfrica)
	fmt.Println(canSeeAsia)
	fmt.Println(canSeeNA)
	fmt.Println(canSeeSA)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadQuarters&roles == isHeadQuarters)
}
