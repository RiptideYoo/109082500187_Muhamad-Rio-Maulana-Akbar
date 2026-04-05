package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	// Membaca 8 nilai waktu pengerjaan
	for i := 1; i <= 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++        
			*skor += waktu 
		}
	}
}
func main(){
    var nama, pemenang string
    var soal, skor, maxSoal, minSkor int

    for {
        fmt.Scan(&nama)
        if nama == "Selesai" {
            break
        }

        hitungSkor(&soal, &skor)
        if soal > maxSoal || (soal == maxSoal && skor < 9999) {
            maxSoal = soal
            minSkor = skor
            pemenang = nama
        }
    }

    fmt.Println(pemenang, maxSoal, minSkor)
}
