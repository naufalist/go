package main

import "fmt"

func main() {
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "JavaScript"
	// languages[3] = "C#"
	// languages[4] = "Python"

	// languages := [5]string{"Ruby", "Python", "JavaScript", "Go", "C#"}

	// languages := [5]string{
	// 	"Ruby",
	// 	"Python",
	// 	"JavaScript",
	// 	"Go",
	// 	"C#",
	// }

	languages := [...]string{
		"Ruby",
		"Python",
		"JavaScript",
		"Go",
		"C#",
	}

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	for index, lang := range languages {
		fmt.Println("Index : ", index, " language : ", lang)
	}
}
