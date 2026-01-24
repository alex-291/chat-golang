package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your age")
	var age int
	itemsScanned, err := fmt.Scan(age)
	fmt.Println("item scanned", itemsScanned)
	fmt.Println(err.Error())
	fmt.Println("Your age", age)
}
