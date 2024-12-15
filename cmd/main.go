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
		// Chapter 001 - Vars and Constants
		examples.Vars_constants_declaration()

	case 2:
		// Chapter 002 - Input Data on CLI
		examples.Input_data_to_vars()
		examples.Other_Input_func()

	case 3:
		// Chapter 003 - Functions
		examples.UsingFunctions()

	case 4:
		// Chapter 004 - Switchs
		examples.UsingSwitch()

	case 5:
		// Chapter 005 - Break, Continue and Defer
		examples.UsingDeferBreakContinue()

	case 6:
		// Chapter 006 - Slices
		examples.UsingArrayAndSlices()

	case 7:
		// Chapter 007 - Maps
		examples.UsingMaps()

	case 8:
		// Chapter 008 - Structs
		examples.UsingStructs()

	case 9:
		// Chapter 009 - Packages
		examples.UsingPackages()

	case 10:
		// Chapter 010 - Simulating POO
		examples.UsingPoo()

	case 11:
		// Chapter 011 - Interfaces
		examples.UsingInterfaces()

	case 12:
		// Chapter 012 - Concurrency 1: Go Rutines
		examples.UsingConcurrentGorutines()

	case 13:
		// Chapter 013 - Concurrency 2: Channels
		examples.UsingConcurrentChannels()
		examples.UsingMoreChannels()

	case 14:
		// Chapter 014 - Webserver External Lib: Echo
		examples.CreatingAnEchoServer()

	case 15:
		// Chapter 015 - Webserver Standard Lib with routing
		fmt.Printf("\nIn Progress...\n")

	default:
		fmt.Printf("\nInvalid Option! Exit Program...\n")
	}
}
