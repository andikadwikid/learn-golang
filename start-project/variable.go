package main

import "fmt"

func main() {
	var nama string

	nama = "Andika DWiki"
	fmt.Println(nama)

	nama = "Andika Dwiki Darmawan"
	fmt.Println(nama)

	// variabel tanpa tipe data
	var umur = 23
	fmt.Println(umur)

	// variabel tanpa deklarasi
	alamat := "Jakarta"
	fmt.Println(alamat)

	// multiple variable
	var (
		namaDepan = "Andika"
		namaBelakang = "Dwiki"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}
