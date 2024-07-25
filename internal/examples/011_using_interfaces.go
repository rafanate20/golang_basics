package examples

import (
	"fmt"

	mpk "github.com/rafanate20/golang_basics/internal/mypackage"
)

func UsingInterfaces() {
	fmt.Println("Chapter 011 - Using a custom pkg with Interfaces:")

	mySquare := mpk.Square{Base: 5}
	fmt.Println(mpk.CalculateArea(mySquare))
	myRectangle := mpk.Rectangle{Width: 4, High: 5}
	fmt.Println(mpk.CalculateArea(myRectangle))
	myTrapezoid := mpk.Trapezoid{BaseA: 18, BaseB: 10, High: 5}
	fmt.Println(mpk.CalculateArea(myTrapezoid))
	myCircle := mpk.Circle{Radio: 4}
	fmt.Println(mpk.CalculateArea(myCircle))

	//Lista de Interfaces
	myInterface := []interface{}{"figure", 1, 5.9, false}
	fmt.Println(myInterface...)
	fmt.Println("----------------------------------------------------------------")
}
