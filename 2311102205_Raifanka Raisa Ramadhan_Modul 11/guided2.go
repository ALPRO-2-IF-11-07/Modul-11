package main

import "fmt"

// Define arrInt as an array of integers with a fixed size
type arrInt [4321]int

// Binary search function for descending sorted array
func BinarySearch_2(T arrInt, n int, X int) int {
	/* Mengembalikan indeks dari X apabila X ditemukan di dalam array T
	   yang berisi n buah bilangan bulat terurut secara descending/menciut,
	   atau -1 apabila X tidak ditemukan */
	var found int = -1
	var med int
	var kr int = 0
	var kn int = n - 1

	// Iterasi mencari elemen dengan binary search
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if X > T[med] { // Karena descending, jika X lebih besar, bergerak ke kiri
			kn = med - 1
		} else if X < T[med] { // Jika X lebih kecil, bergerak ke kanan
			kr = med + 1
		} else { // Jika ditemukan
			found = med
		}
	}
	return found
}

func main() {
	// Contoh penggunaan
	var data arrInt
	var n int
	// Input jumlah elemen
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scanln(&n)
	// Input elemen array (harus terurut secara descending)
	fmt.Println("Masukkan elemen array (harus terurut menurun):")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i+1)
		fmt.Scanln(&data[i])
	}

	// Input elemen yang dicari
	var search int
	fmt.Print("Masukkan elemen yang ingin dicari: ")
	fmt.Scanln(&search)

	// Panggil fungsi binary search
	result := BinarySearch_2(data, n, search)

	// Cetak hasil
	if result == -1 {
		fmt.Println("Elemen tidak ditemukan dalam array.")
	} else {
		fmt.Printf("Elemen ditemukan pada indeks: %d\n", result)
	}
}
