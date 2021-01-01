package main

import "fmt"

func main() {
	myMap := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"JavaScript": "hype",
	}

	value, isAvailable := myMap["python"]
	fmt.Println(isAvailable)
	fmt.Println(value)

	// fmt.Println(myMap)
	for key, value := range myMap {
		fmt.Println("Key : ", key, " Value : ", value)
	}

	fmt.Println("==========")

	delete(myMap, "ruby")

	for key, value := range myMap {
		fmt.Println("Key : ", key, " Value : ", value)
	}

	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["ruby"] = 9
	// myMap["go"] = 9
	// myMap["JavaScript"] = 8
	// myMap["go"] = 10

	// fmt.Println(myMap["ruby"])
	// fmt.Println(myMap)
}
