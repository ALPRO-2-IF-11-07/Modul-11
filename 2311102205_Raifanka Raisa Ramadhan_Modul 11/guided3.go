package main

import "fmt"

// Define arrStr as an array of strings with a fixed size
type arrStr [1234]string

// Sequential search function
func SeqSearch_1(T arrStr, n int, X string) int {
	/* Mengembalikan indeks dari X apabila X ditemukan di dalam array T
	   yang berisi n buah teks, atau -1 apabila X tidak ditemukan */
	var found int = -1
	var j int = 0

	// Iterasi mencari elemen yang cocok
	for j < n && found == -1 {
		if T[j] == X {
			found = j
		}
		j = j + 1
	}
	return found
}
func main() {
	// Contoh penggunaan
	var data arrStr
	var n int

	// Input jumlah elemen
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scanln(&n)

	// Input elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i+1)
		fmt.Scanln(&data[i])
	}
	// Input elemen yang dicari
	var search string
	fmt.Print("Masukkan elemen yang ingin dicari: ")
	fmt.Scanln(&search)
	// Panggil fungsi sequential search
	result := SeqSearch_1(data, n, search)
	// Cetak hasil
	if result == -1 {
		fmt.Println("Elemen tidak ditemukan dalam array.")
	} else {
		fmt.Printf("Elemen ditemukan pada indeks: %d\n", result)
	}
}
