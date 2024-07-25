package examples

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Input_data_to_vars() {
	fmt.Printf("\nChapter 002 - Example of Input Data to Variables:\n")

	//scanner to read console input
	scanner := bufio.NewScanner(os.Stdin)

	//square area
	fmt.Println("Insert Square Base (int):")
	scanner.Scan()
	var baseSquare int
	baseSquare, _ = strconv.Atoi(scanner.Text())
	squareArea := baseSquare * baseSquare
	fmt.Println("Square Area:", squareArea)

	x := 0
	y := 0

	fmt.Println("Insert X (int):")
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())

	fmt.Println("Insert Y (int):")
	scanner.Scan()
	y, _ = strconv.Atoi(scanner.Text())

	//sum
	result := x + y
	fmt.Println("Sum:", result)

	//Subtract
	result = x - y
	fmt.Println("Subtract:", result)

	//multiplication
	result = x * y
	fmt.Println("Multiplication:", result)

	//Division
	result = x / y
	fmt.Println("Division:", result)

	//Modulus
	result = x % y
	fmt.Println("Modulus:", result)

	//Incremental - Decremental
	x++
	y--
	fmt.Println("Incremental:", x, "Decremental:", y)

	fmt.Println("----------------------------------------------------------------")
}

func Other_Input_func() {
	//scanner to read console input
	scanner := bufio.NewScanner(os.Stdin)

	//vars
	company := "MyCompany"
	old := 6

	//getting company name
	fmt.Println("Insert Name (String):")
	scanner.Scan()
	company = scanner.Text()

	//getting years old
	fmt.Println("Insert Old (Int):")
	scanner.Scan()
	old, err := strconv.Atoi(scanner.Text())

	if err == nil {
		//Printf
		fmt.Printf("%s is %d years old.\n", company, old)
		fmt.Printf("%v is %v years old.\n", company, old)

		//Sprintf
		msg := fmt.Sprintf("%s is %d years old.", company, old)
		fmt.Println(msg)

		//Var Type
		fmt.Printf("company is %T and old is %T\n", company, old)
	} else {
		fmt.Println("Error converting Int value")
	}

	fmt.Println("----------------------------------------------------------------")
}
