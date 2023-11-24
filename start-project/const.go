package main

import "fmt"

func main() {
	const nama string = "Andika Dwiki"

	fmt.Println(nama)

	const (
		namaDepan    = "Andika"
		namaBelakang = "Dwiki"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}
