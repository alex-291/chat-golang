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

	word := "New"
	fmt.Println(word)

	// make([]тип, сколько_пассажиров, сколько_кресел)
	bus := make([]int, 0, 5)

	fmt.Println("Сколько пассажиров (len):", len(bus))
	fmt.Println("Сколько всего кресел (cap):", cap(bus))

	bus = append(bus, 10) // посадили одного
	bus = append(bus, 20) // посадили второго

	fmt.Println("Пассажиров:", len(bus)) // теперь 2
	fmt.Println("Кресел:", cap(bus))     // всё еще 5s

	// Допустим, мы заполнили все 5 мест
	bus = append(bus, 30, 40, 50)
	fmt.Println("Заполнено:", len(bus), "из", cap(bus)) // 5 из 5

	// Пытаемся посадить 6-го!
	bus = append(bus, 60)

	fmt.Println("После перебора:")
	fmt.Println("Новый Len:", len(bus)) // станет 6
	fmt.Println("Новый Cap:", cap(bus)) // станет 10! (удвоился)

}
