package main

import (
	"chat/src/data"
)

func main() {
	alex := data.User{
		Name: "Alex",
		Age:  54,
	}
	valera := data.User{
		Name: "Valera",
		Age:  21,
	}

	data.PrintUserInfo(alex, valera)

}
