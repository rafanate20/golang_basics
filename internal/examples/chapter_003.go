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
}
