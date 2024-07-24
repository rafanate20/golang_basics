package examples

import "fmt"

func useDefer() {
	fmt.Println("Executed at the end using defer")
	fmt.Println("\n----------------------------------------------------------------")
}

func useContinue() {
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
		if i > 2 && i <= 5 {
			fmt.Print("\n")
			continue
		}
		fmt.Printf(" Executing more lines...\n")
	}
}

func useBreak() {
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
		if i > 2 && i <= 5 {
			fmt.Print("\n")
			break
		}
		fmt.Printf(" Executing more lines...\n")
	}
}

func UsingDeferBreakContinue() {
	defer useDefer()

	fmt.Println("Calling function with continue")
	useContinue()

	fmt.Println("Calling function with break")
	useBreak()
}
