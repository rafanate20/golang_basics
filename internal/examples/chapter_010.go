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

	// Class composition and use interfaces to simulate Polymorphism
	ftEmp := mpk.FullTimeEmployee{}
	ftEmp.Name = "Juan Guarnizo"
	ftEmp.Birthdate = "25/02/1992"
	ftEmp.Id = 2566
	ftEmp.Position = "streamer"
	ftEmp.ContractId = 55456
	ftEmp.GetBirthdayMsg()
	mpk.GetMsg(ftEmp)

	ptEmp := mpk.PartTimeEmployee{}
	ptEmp.Name = "Fernando Flow"
	ptEmp.Birthdate = "10/08/2001"
	ptEmp.Id = 2463
	ptEmp.Position = "housekeeper"
	ptEmp.HoursPerWeek = 12
	ptEmp.GetBirthdayMsg()
	mpk.GetMsg(ptEmp)
	fmt.Println("----------------------------------------------------------------")

	// Class factory to use the common desing pattern Abstract Factory
	smsFactory, _ := mpk.GetNotificationFactory("SMS")
	emailFactory, _ := mpk.GetNotificationFactory("Email")

	// call SMS factory functions
	mpk.GetMethod(smsFactory)
	mpk.GetChannel(smsFactory)
	mpk.SendNotification(smsFactory)

	// call eMail factory functions
	mpk.GetMethod(emailFactory)
	mpk.GetChannel(emailFactory)
	mpk.SendNotification(emailFactory)
}
