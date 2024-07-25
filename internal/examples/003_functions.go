package examples

import (
	"fmt"
	"log"
	"strconv"
)

// Functions started with lowercase letter is not exported
func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

// Function that returns values by their name.
func getValues(x int) (double int, triple int, quad int) {
	// return 2 * x, 3 * x, 4 * x // old way.
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

/*
VARIADIC Functions
when the number of elements to receive can be variable,
it is possible to create a variadic function that
can be called with any number of trailing arguments.
*/
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}

	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func UsingFunctions() {
	fmt.Printf("\nChapter 003 - Using not exposed Functions :\n")

	//function returning 2 values
	val1, val2 := doubleReturn(5)
	fmt.Println(val1, val2)

	//receiving just 1 value
	_, justSecVal := doubleReturn(15)
	fmt.Println(justSecVal)

	//Converting a Value
	convertedValue, error := strconv.Atoi("67657")
	if error != nil {
		log.Fatal(error)
	} else {
		println("Converted Value: ", convertedValue)
	}

	fmt.Println("----------------------------------------------------------------")

	// print multiple returns:
	fmt.Println(getValues(5))

	// receiving multiple returns on new vars:
	d, t, q := getValues(5)
	fmt.Println(d, t, q)

	fmt.Println("----------------------------------------------------------------")

	// calling a variadic function with multiple int variables.
	fmt.Println(sum(3, 5, 7, 10, 25, 4))

	// calling the variadic function with a slice
	numbers := []int{4, 7, 3}
	fmt.Println(sum(numbers...))

	// calling a variadic function with strings
	printNames("Carlos", "Jose", "Maria", "Laura")

}
