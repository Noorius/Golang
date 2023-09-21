package main

import (
	"HW1/accountant"
	"HW1/designer"
	"HW1/manager"
	"fmt"
)

func main() {
	acc := accountant.Accountant{}
	acc.SetSalary(100.0)
	acc.SetPosition("Accountant")
	acc.SetAddress("Almaty")
	fmt.Println("I am", acc.GetPosition(), "getting", acc.GetSalary(), "and living in", acc.GetAddress())

	mng := manager.Manager{}
	mng.SetSalary(200)
	mng.SetPosition("Manager")
	mng.SetAddress("Astana")
	fmt.Println("I am", mng.GetPosition(), "getting", mng.GetSalary(), "and living in", mng.GetAddress())

	des := designer.Designer{}
	des.SetSalary(100)
	des.SetAddress("Atyrau")
	des.SetPosition("Designer")
	fmt.Println("I am", des.GetPosition(), "getting", des.GetSalary(), "and living in", des.GetAddress())
}
