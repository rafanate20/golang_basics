package package_poo

import "fmt"

type Person struct {
	Name      string
	Birthdate string
}

// Person functions
func (p Person) GetBirthdayMsg() {
	fmt.Printf("Hi, my name is %s, and I was born in %s. ", p.Name, p.Birthdate)
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	// Class Composition
	Person
	Employee
	ContractId int
}

type PartTimeEmployee struct {
	// Class Composition
	Person
	Employee
	HoursPerWeek int
}

type EmployeeFuncs interface {
	GetEmployeeMsg() string
}

// Employee functions per type (acording to the EmployeeFuncs interface)
func (ftemp FullTimeEmployee) GetEmployeeMsg() string {
	return "a Full Time Employee"
}

func (ptemp PartTimeEmployee) GetEmployeeMsg() string {

	return "a Part Time Employee"
}

// Static functions (Implementing the EmployeeFuncs interface)
func GetMsg(iemp EmployeeFuncs) {
	fmt.Printf("I am %s\n", iemp.GetEmployeeMsg())
}
