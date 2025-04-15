package main

import "fmt"

func main() {

	myNumber := 5

	var ptr = &myNumber

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 3
	fmt.Println(myNumber)

}
