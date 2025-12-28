package main

import "fmt"

func main(){
	for count := range 20{
		fmt.Println("Hello world", count)
	}
}