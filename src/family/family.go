package family

import "fmt"

func GreetUser(username string) {
	fmt.Println("Hi!", username, "We are family")
}

func ByeByeUser(username string, date int, month string) {
	fmt.Println("Goodbye", username, "see you", date, month)
}
