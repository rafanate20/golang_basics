package examples

import "fmt"

func UsingSwitch() {
	fmt.Printf("\nChapter 004 - Using Switch :\n")

	//Value
	val := 102

	//Clasicc Switch
	switch mod := val % 2; mod {
	case 0:
		fmt.Println("It's pair")
	default:
		fmt.Println("It isn't pair")
	}

	//Go Switch to replace to much ifs sentences
	switch {
	case val > 100:
		fmt.Println("It's bigger")
	case val <= 0:
		fmt.Println("It's lower")
	default:
		fmt.Println("It's in the middle")
	}

	fmt.Println("----------------------------------------------------------------")
}
