package main

import (
	"chat/alex"
	"chat/nat"
	"chat/user"
	"chat/val"
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
}
