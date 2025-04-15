package main

import "fmt"

func main() {
	fmt.Println("Array -->we dont use much in Golang unlike other languages")

	var fruits [3]string

	fruits[0] = "Mango"
	//fruits[1] = "Grapes"
	fruits[2] = "Guava"
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var vegetables = [3]string{"Mushroom", "Beetroot", "Pumpkin"}
	fmt.Println(vegetables)

}
