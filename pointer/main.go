package main

import "fmt"

func main() {
	// #1
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 10

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	// #2
	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 20

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// #3
	number := 5
	fmt.Println("Alamat memory :", &number)
	fmt.Println("Nilai awal :", number)
	change(&number, 100)
	fmt.Println("Nilai akhir :", number)
	fmt.Println("Alamat memory :", &number)
}

func change(old *int, new int) {
	*old = new

	fmt.Println("Alamat memory :", old)
}
