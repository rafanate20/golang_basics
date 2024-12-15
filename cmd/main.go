package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/rafanate20/golang_basics/internal/examples"
)

func main() {
	fmt.Printf(`
	......... GoLang Basics Examples ........
	.........................................
	 1.- Variables and Constants
	 2.- Console Input Variables
	 3.- Functions
	 4.- Switchs
	 5.- Defer, Break and Continue
	 6.- Arrays and Slices
	 7.- Maps
	 8.- Structs
	 9.- Packages
	10.- POO
	11.- Interfaces
	12.- GoRutines (concurrency)
	13.- Channels (concurrency)
	14.- WebServer (Using Echo Framework)
	15.- WebServer (stdlib)
	.........................................
	INSERT an option number:
	`)
	//scanner to read console input
	scanVar := bufio.NewScanner(os.Stdin)
	scanVar.Scan()

	optionSelected, optErr := strconv.Atoi(scanVar.Text())
	if optErr != nil {
		fmt.Printf("\nInvalid Option! Exit Program...\n")
	}

	switch optionSelected {
	case 1:
		// Chapter 001
		examples.Vars_constants_declaration()

	case 2:
		// Chapter 002
		examples.Input_data_to_vars()
		examples.Other_Input_func()

	case 3:
		// Chapter 003
		examples.UsingFunctions()

	case 4:
		// Chapter 004
		examples.UsingSwitch()

	case 5:
		// Chapter 005
		examples.UsingDeferBreakContinue()

	case 6:
		// Chapter 006
		examples.UsingArrayAndSlices()

	case 7:
		// Chapter 007
		examples.UsingMaps()

	case 8:
		// Chapter 008
		examples.UsingStructs()

	case 9:
		// Chapter 009
		examples.UsingPackages()

	case 10:
		// Chapter 010
		examples.UsingPoo()

	case 11:
		// Chapter 011
		examples.UsingInterfaces()

	case 12:
		// Chapter 012
		examples.UsingConcurrentGorutines()

	case 13:
		// Chapter 013
		examples.UsingConcurrentChannels()
		examples.UsingMoreChannels()

	case 14:
		// Chapter 014
		examples.CreatingAnEchoServer()

	default:
		fmt.Printf("\nInvalid Option! Exit Program...\n")
	}
}
