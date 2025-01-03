package package_poo

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func New(rm, dsk int, brn string) Pc {
	myPc := Pc{ram: rm, disk: dsk, brand: brn}
	return myPc
}

func (myPc Pc) FormatPrint() string {
	return fmt.Sprintf("Esta PC marca %s cuenta con una ram de %dGB y un disco de %dGB.", myPc.brand, myPc.ram, myPc.disk)
}

func (myPc *Pc) DuplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func (myPc Pc) GetRam() int {
	return myPc.ram
}

func (myPc *Pc) SetRam(rm int) {
	myPc.ram = rm
}
