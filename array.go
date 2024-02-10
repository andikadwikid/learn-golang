package main

import "fmt"

func main() {
	//buat variabel dan tentukan panjang array, lalu tentukan juga tipe data array
	//disini panjang array nya adalah 3 dengan tipe data string [3]string
	var names [3]string

	names[0] = "Andika"
	names[1] = "Dwiki"
	names[2] = "Darmawan"

	fmt.Println(names[0])

	//bisa juga dideklarasikan langsung
	var values = [3]int{90, 80, 100}
	fmt.Println(values)

	//bisa juga panjang array ditentukan, menggunakan [...]
	var values2 = [...]int{90, 80, 100, 150}
	fmt.Println(values2)

	//mencari panjang array
	fmt.Println(len(names))

	//mencari posisi index array
	fmt.Println(names[0:2])

	//mengubah posisi index array
	names[0] = "Andika Dwiki"

	fmt.Println(names)
}
