package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Naufal", "score": "A"},
		{"name": "Link", "score": "B"},
		{"name": "Steven", "score": "E"},
	}

	for _, student := range students {
		fmt.Println(student["score"])
	}

	// fmt.Println(students)
}
