package main

import "fmt"

func main() {

	score := 65
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)

	// age := 9

	// if age > 10 {
	// 	fmt.Println("Boleh bermain game")
	// } else {
	// 	fmt.Println("Kamu belum boleh")
	// }

}

//if
//if else
//else if
//if bersarang