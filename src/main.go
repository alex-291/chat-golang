package main

import "fmt"

func main() {
	number := uint8(255)

	fmt.Println("value", number)
	fmt.Println("address", &number)

	numberAddress := &number
	fmt.Println(numberAddress)

	numbersArray := [3]int{10, 20, 30}
	fmt.Println(numbersArray, "length:", len(numbersArray))
	numbersArray[1] = 7
	for i, v := range numbersArray {
		fmt.Println("На полке №", i, "лежит:", v)
	}

	numbersSlice := []int{10}
	fmt.Println(numbersSlice, "length:", len(numbersSlice), "capacity:", cap(numbersSlice))

	numbersSlice = append(numbersSlice, -2)
	fmt.Println(numbersSlice, "length:", len(numbersSlice), "capacity:", cap(numbersSlice))

	firstElement := numbersSlice[0]
	fmt.Println(firstElement)

	numbersSlice[0] = 46
	fmt.Println(numbersSlice, "length:", len(numbersSlice), "capacity:", cap(numbersSlice))

}
