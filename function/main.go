package main

import "fmt"

func main() {
	luas, _ := calculate(10, 2)
	fmt.Println(luas)
	// fmt.Println(keliling)

	// sentence := printMyResult("hei")
	// fmt.Println(sentence)
	// printMyResult("kamu")
	// printMyResult("iya, kamu")
}

func printMyResult(sentence string) string {
	// fmt.Println(sentence)
	newSentence := sentence + " kamu"
	return newSentence
}

func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}
