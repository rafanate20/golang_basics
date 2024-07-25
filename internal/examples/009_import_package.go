package examples

import (
	"fmt"

	mpk "github.com/rafanate20/golang_basics/internal/mypackage"
)

func UsingPackages() {
	fmt.Println("Chapter 009 - Using Packages:")

	var myCar mpk.VehiclePublic
	myCar.Brand = "GM"
	myCar.Model = "Hummer"
	myCar.Year = 2020
	//myCar.stockid //is private

	fmt.Println("before generate private id", myCar)

	myCar = mpk.GenerateStockid(myCar)
	fmt.Println("after generate private id", myCar)

	fmt.Println("----------------------------------------------------------------")
}
