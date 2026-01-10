package nat

import "fmt"

func GreetUser(name string, age int) {
	fmt.Println("This is massage from package nat")
	fmt.Println("Hello, everybody! I'm", name)
	fmt.Println("I'm", age, "old years")
}
