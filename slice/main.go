package main

import "fmt"

func main() {
	// jenisLayangan := []string{"Sukhoi", "Seot", "Munaroh"}

	var jenisLayangan []string
	jenisLayangan = append(jenisLayangan, "Sukhoi")
	jenisLayangan = append(jenisLayangan, "Seot")
	jenisLayangan = append(jenisLayangan, "Munaroh")

	// fmt.Println(jenisLayangan)

	for _, layangan := range jenisLayangan {
		fmt.Println(layangan)
	}
}
