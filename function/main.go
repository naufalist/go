package main

import "fmt"

func main() {
	sentence := printMyResult("hei")
	fmt.Println(sentence)
	// printMyResult("kamu")
	// printMyResult("iya, kamu")
}

func printMyResult(sentence string) string {
	// fmt.Println(sentence)
	newSentence := sentence + " kamu"
	return newSentence
}
