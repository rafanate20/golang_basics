package examples

import "fmt"

type vehicle struct {
	brand string
	model string
	year  uint32
}

func UsingStructs() {
	fmt.Println("Chapter 008 - Using Structs:")

	myCar := vehicle{brand: "Nissan", model: "Versa", year: 2018}
	fmt.Println("My Car", myCar)

	var herCar vehicle
	herCar.brand = "Mazda"
	herCar.model = "3"
	fmt.Println("Her Car", herCar)

	fmt.Println("----------------------------------------------------------------")
}
