package examples

import (
	"fmt"

	mpk "github.com/rafanate20/golang_basics/internal/package_poo"
)

func UsingPoo() {
	fmt.Println("Chapter 010 - Using POO with Structs and Pointers:")

	myPc := mpk.New(12, 200, "HP")
	myPc.SetRam(16)
	fmt.Println(myPc.FormatPrint())
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	fmt.Println(myPc.FormatPrint())
	fmt.Println("----------------------------------------------------------------")
}
