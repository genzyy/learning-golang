package main

import "fmt"

func main() {
	var n bool = true
	m := 1 == 1
	fmt.Printf("%v %T\n", n, n)
	//fmt.Println()
	fmt.Printf("%v %T\n", m, m)
	// Every boolean variable without any value is by default set to false.

	// Unsigned integer type
	var j uint = 40
	fmt.Printf("%v %T\n", j, j)
	fmt.Println(j)

	// Mismatched type cant take part in an expression.
	var a int = 3
	var b int8 = 2
	//fmt.Print(a / b)
	// To fix the type mismatch we have to convert
	// one of the variables to make the expression which has only one type.
	fmt.Print(a / int(b))

	// Primitives with binary numbers
	c := 10
	d := 2
	fmt.Println(c & d)
	fmt.Println(c | d)
	fmt.Println(c ^ d)
	fmt.Println(c &^ d)
	// Shift Operations
	// What these shift operator does is first convert the given number
	// into binary and then take 1 and shift that 1 in the operator direction.
	fmt.Println(c << 3)
	fmt.Println(c >> 3)

	// Floating type
	var pi float32 = 3.14
	//pi = 13.7e72 Too big for float32 but fine for float64
	pi = 2.1e14
	fmt.Printf("%v %T\n", pi, pi)

	// Complex Numbers
	var iota complex64 = 1 + 2i
	fmt.Printf("%v %T\n", iota, iota)
	fmt.Printf("%v %T\n", real(iota), real(iota))
	fmt.Printf("%v %T\n", imag(iota), imag(iota))

	var num complex128 = complex(5, 12)
	fmt.Println(num)

	var str string = "This is also a string"
	fmt.Println(str)
	// We cannot change the string like this because its type is byte => immutable
	//str[2] = "u"
	fmt.Printf("%v %T", str[2], str[2])
	// String Concatenation.
	fmt.Println(str + "This is a string")

	var sentence string = "This is a sentence"
	numsen := []byte(sentence)
	fmt.Printf("%v %T\n", sentence, sentence)
	fmt.Printf("%v %T\n", numsen, numsen)
	// The string type represents utf8 characters

	// Rune type
	// The rune type represents utf32 characters.
	var ru rune = 'a'
	fmt.Printf("%v %T\n", ru, ru)
	var r rune = 'b'
	fmt.Printf("%v %T\n", r, r)
}
