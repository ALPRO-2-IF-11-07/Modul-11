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

	for calon := 1; calon <= 20; calon++ {
		if suaraCalon[calon] > 0 {
			fmt.Printf("%d: %d\n", calon, suaraCalon[calon])
		}
	}
}
