package main

import (
	"fmt"
)

func main() {
	var suara int
	var totalSuaraMasuk, totalSuaraSah int

	suaraCalon := [21]int{}

	for {
		fmt.Scan(&suara)

		if suara == 0 {
			break
		}

		totalSuaraMasuk++

		if suara >= 1 && suara <= 20 {
			totalSuaraSah++

			suaraCalon[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)

	var ketua, wakil int
	var suaraTerbanyak, suaraWakil int

	for calon := 1; calon <= 20; calon++ {
		if suaraCalon[calon] > suaraTerbanyak ||
			(suaraCalon[calon] == suaraTerbanyak && calon < ketua) {
			wakil = ketua
			suaraWakil = suaraTerbanyak
			ketua = calon
			suaraTerbanyak = suaraCalon[calon]
		} else if suaraCalon[calon] > suaraWakil ||
			(suaraCalon[calon] == suaraWakil && calon < wakil) {
			wakil = calon
			suaraWakil = suaraCalon[calon]
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
