package main

import (
	"chat/src/alex"
	"chat/src/nat"
	"chat/src/val"
	"fmt"
)

func main() {
	age, name := 51, "alex"
	fmt.Println(age, name)

	anotherage, anothername := alex.GetAgeAndName()
	fmt.Println(anotherage, anothername)

	address := alex.GetAddress()
	fmt.Println(address)

	ageN, nameN := 55, "Nat"
	fmt.Println(ageN, nameN)

	onemoreage, onemorename := nat.GetAgeAndName()
	fmt.Println(onemoreage, onemorename)

	onemoreaddress := nat.GetAddress()
	fmt.Println(onemoreaddress)

	nameV, ageV := "Valery", 21
	fmt.Println(nameV, ageV)

	anothernameV, anotherageV := val.GetAgeAndName()
	fmt.Println(anothernameV, anotherageV)

	anotheraddressV := val.GetAddress()
	fmt.Println(anotheraddressV)
}
