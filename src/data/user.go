package data

import "fmt"

type User struct {
	Name string
	Age  uint8
}

func GreetUser(userName string) {
	fmt.Println("Hello", userName)
}

func PrintUserInfo(users ...User) {
	for _, user := range users {
		address := fmt.Sprintf("%s's address: %p", user.Name, &user)
		fmt.Println(user, address)
	}
}
