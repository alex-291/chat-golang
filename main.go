package main

import "fmt"

func main() {
	fmt.Println("func main was called")
	printPersonalData()
	acceptor("Alex", "Timofeev", 54)
	fmt.Println("func main has ended")
}

func printPersonalData() {
	fmt.Println("func printPersonalData was called")
	firstName := "Alex"
	secondName := "Timofeev"
	age := 54
	fmt.Println(firstName)
	fmt.Println(secondName)
	fmt.Println(age)
	fmt.Println("func printPersonalData has ended")
}

func acceptor(
	firstName string,
	secondName string,
	age int,
) {
	fmt.Println("func acceptor was called")
	fmt.Println(firstName, secondName, age)
	fmt.Println("func acceptor hs ended")

}
