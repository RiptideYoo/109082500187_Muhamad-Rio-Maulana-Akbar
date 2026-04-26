package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nIsi array:")
	fmt.Println(arr)

	fmt.Println("\nElemen indeks ganjil:")
	for i := 0; i < N; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}

	fmt.Println("\nElemen indeks genap:")
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	var x int
	fmt.Print("\nMasukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Println("Elemen indeks kelipatan", x, ":")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	var idx int
	fmt.Println("\nmasukkan index yang mau dihapus: ")
	fmt.Scan(&idx)
	temp := arr[:idx]  
	
	for i := idx + 1; i < len(arr); i++ {
    temp = append(temp, arr[i])
	}
	arr = temp

	fmt.Println("Array setelah penghapusan:")
	fmt.Println(arr)

	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	rata := float64(total) / float64(len(arr))
	fmt.Println("Rata-rata:", rata)

	var jumlah float64
	for i := 0; i < len(arr); i++ {
		selisih := float64(arr[i]) - rata
		jumlah += selisih * selisih
	}
	stdDev := math.Sqrt(jumlah / float64(len(arr)))
	fmt.Println("Standar deviasi:", stdDev)

	var cari int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)

	frekuensi := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == cari {
			frekuensi++
		}
	}

	fmt.Println("Frekuensi", cari, "=", frekuensi)
}