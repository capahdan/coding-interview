package pinjaman

import "fmt"

func Pinjaman(pokok, bunga, float64, tenor int) {
	pinjamanPerBulan := (pokok / tenor)
	bungaPerBulan := (pokok / tenor) * bunga

	totalPinjaman := pokok + (bunga * tenor * pokok)

	fmt.Println(pinjamanPerBulan, bungaPerBulan, totalPinjaman)
	for i := 1; i <= tenor; i++ {
		fmt.Println(i)
	}

}
