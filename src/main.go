package main

import (
	"chat/src/alex"
	"chat/src/family"
	"chat/src/nat"
	"chat/src/user"
	"chat/src/val"
)

func main() {
	name := "Max"
	user.GreetUser(name)

	age := 54
	alex.PrintAge(age)

	name = "Nat"
	nat.GreetUser(name, age)

	age = 21
	val.AboutMe(age)

	anorherName := "Cris"
	family.GreetUser(anorherName)

	date := 29
	month := "February"
	family.ByeByeUser(anorherName, date, month)
}
