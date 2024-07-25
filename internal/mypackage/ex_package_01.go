package mypackage

import (
	"fmt"
	"math/rand"
)

type VehiclePublic struct { //Begin Uppercase is Public
	Brand   string
	Model   string
	Year    uint32
	stockid int //private when you declare starting with a lower case
}

type contractPrivate struct { //Begin lowercase is private
	id     int
	vendor string
}

// setContract is a private func
func setContract() int {
	var nContract contractPrivate
	nContract.id = rand.Intn(99999)
	nContract.vendor = func(n string) string { // using an anonymous function just for fun
		return n
	}("Juan")
	fmt.Printf("the vendor %s set a contract number %d", nContract.vendor, nContract.id)

	return nContract.id
}

// GenerateStockid is a public func
func GenerateStockid(car VehiclePublic) VehiclePublic {
	car.stockid = setContract()

	return car
}
