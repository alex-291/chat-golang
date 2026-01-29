package main

import "fmt"

func main() {

	var otherdate int

	otherdate = 101

	fmt.Println(otherdate)

	var name string

	name = "This is test, This is test, v, This is test"

	fmt.Println(name)

	var byteArray [10]byte

	fmt.Println(byteArray)

	var intArray [3]int

	intArray = [3]int{299, 10, 50}

	fmt.Println(intArray, "length:", len(intArray))

	fmt.Println("first number is", intArray[2])

	boolArray := [1]bool{true}

	fmt.Println(boolArray)

	number := [3]int{10, 20, 30}

	fmt.Println(number)

	fmt.Println(number, "length:", len(number))

	number[1] = 7

	for i, v := range number {
		fmt.Println("На полке №", i, "лежит:", v)
	}
}
