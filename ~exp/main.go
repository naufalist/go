package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	fmt.Println("--- test reader ---")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input text: ")
	text, _ := reader.ReadString('\n')
	// fmt.Println("This is your text:", text)
	fmt.Printf("This is your text: %s", text)

}
