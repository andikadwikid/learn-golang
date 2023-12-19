package main

import "fmt"

func main() {
	type NoKTP string

	var ktpDika NoKTP = "123456789"

	var contoh string = "123"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpDika)
	fmt.Println(contohKTP)

}
