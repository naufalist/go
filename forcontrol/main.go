package main

import "fmt"

func main() {

	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i, "Saya sedang belajar Go")
	// }

	// i := 1
	// for i < 10 {
	// 	fmt.Println(i, "Saya sedang belajar Go")
	// 	i++
	// }

	title := "Golang the best language"

	// for index, letter := range title {
	// 	fmt.Println("index: ", index, " letter :", string(letter))
	// }

	for _, letter := range title {
		fmt.Println(" letter :", string(letter))
	}
}