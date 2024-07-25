package examples

import "fmt"

func Vars_constants_declaration() {
	fmt.Printf("\nChapter 001 - Example of Variables and Constants Declaration:\n")

	//Constants
	const pi float32 = 3.14
	const pi2 = 3.1416

	//Console Prints
	fmt.Println("PI:", pi)
	fmt.Println("PI2:", pi2)

	//Variable declaration
	base := 12
	var high int = 14
	var area int

	fmt.Println("Vars:", base, high, area)

	//Zero values - AutoAssing
	var a int
	var b float32
	var c string
	var d bool

	fmt.Println("Zero Values:", a, b, c, d)

	//square area
	const baseSquare = 10
	squareArea := baseSquare * baseSquare
	fmt.Println("Square Area:", squareArea)

	fmt.Println("----------------------------------------------------------------")
}
