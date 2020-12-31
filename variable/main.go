package main

import "fmt"

func main() {

	/*
		string, int, float64, bool

		string: "halo", "golang"
		int: 1, 8888, -1
		float64: 3.14, 100.1
		bool: true, false
	*/

	// var name string = "halo"
	name := "Naufal" // sama dengan -> var name string = "Naufal"
	age := 21        // sama dengan -> var age int = 20
	fmt.Println(name)
	fmt.Println(age)

	age = 20 // ubah nilai gunakan =
	fmt.Println(age)

	// var nim string
	// fmt.Println(nim) // print default value (empty for string, 0 for int)
	// nim = "J3D118042"
	// fmt.Println(nim) // print assigned value

	// var nim string = "J3D118042"
	// fmt.Println(nim)
}