# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>
<p align="center">Muhamad Rio Maulana Akbar - 109082500187</p>

## Unguided 

### 1. 
Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila
diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan
koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan
radiusnya.
#### soal1.go

```go
package main
import (
	"fmt"
	"math"
)
type titik struct{
	x, y int
}
type lingkaran struct{
	pusat titik
	r int
}
func jarak(p, q titik)float64{
	x := (p.x - q.x)*(p.x - q.x)
	y := (p.y - q.y)*(p.y - q.y)
	return math.Sqrt(float64(x)+float64(y))
}
func didalam(c lingkaran, p titik)bool{
	return jarak(c.pusat, p) < float64(c.r)
}
func main() {
	var l1, l2 lingkaran
	var d titik
	var Circle1, Circle2 bool

	fmt.Scanln(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scanln(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&d.x, &d.y)

	Circle1 = didalam(l1, d)
	Circle2 = didalam(l2, d)

	if Circle1 && Circle2 {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	}else if Circle1{
		fmt.Print("Titik di dalam lingkaran 1")
	}else if Circle2{
		fmt.Print("Titik di dalam lingkaran 2")
	}else{
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}



}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%204/output/soal1.png)
Jadi, program ini dibuat untuk menetahui apakah titik lingkaran tersebut termasuk kedalam lingkaran 1 atau 2. cara kerja program ini yaitu pertama diminta imputan berupa titik pusat dan jari-jari dari lingkaran 1 dan 2 kemudian program akan memasukkan inputan tadi kedalam function 'didalam', didalam function tersebut akan dieksekusi apakah jarak dari kedua titik pusat lingkaran kurang dari jari-jarinya. kemudian program akan print dari hasil eksekusi function didalam jika inputan lingakaran 1 'true' maka output nya Titik di dalam lingkaran 1'. 

## Unguided 

### 2. 
Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program
yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array
memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat
menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah
genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh
dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.
Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
tersebut.
h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi
tersebut.
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%204/output/soal2.png)
Jadi, program ini dibuat untuk mengolah sekumpulan bilangan bulat yang disimpan array dengan ukuran tetap sebesar N elemen yang ditentukan oleh pengguna. Cara kerja program ini yaitu pertama pengguna diminta menginput jumlah elemen dan nilai-nilainya, di mana data tersebut disimpan menggunakan fungsi make agar memori yang tersedia pas sesuai kebutuhan. Setelah data tersimpan, program akan mengeksekusi perintah untuk menampilkan isi array keseluruhan, memisahkan berdasarkan indeks ganjil, genap, maupun kelipatan dari angka x yang diinput pengguna. Untuk penghapusan index, program akan mengambil potongan slice dari awal hingga sebelum indeks yang dipilih, lalu menggunakan fungsi append untuk memasukkan sisa elemennya satu per satu secara manual agar data yang dihapus tidak muncul lagi. kemudian mencari nilai rata-rata dan standar deviasi menggunakan rumus akar kuadrat dari math.Sqrt dan menghitung frekuensinya.

## Unguided 

### 3. 
Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang
memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang
digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian
program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan
dalam array adalah nama-nama klub yang menang saja.
Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir
program, tampilkan daftar klub yang memenangkan pertandingan.

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string 

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	i := 1
	for {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		
		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		i++
	}

	for j := 0; j < len(pemenang);j++ {
		fmt.Printf("hasil %d: %s\n", j+1, pemenang[j])
	}
	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%204/output/soal2.png)
Jadi, program ini dibuat untuk pemenang pertandingan bola antara dua klub pada setiap pertandingan . Cara kerja program ini yaitu pertama kita diminta menginput nama Klub A dan Klub B, kemudian program akan masuk ke dalam perulangan untuk meminta masukan skor hasil pertandingan secara terus-menerus. Program akan mengeksekusi logika pengecekan skor: jika skor salah satu klub lebih besar, maka nama klub pemenang tersebut dimasukkan ke dalam wadah slice menggunakan fungsi append, namun jika skornya sama, maka yang disimpan adalah keterangan "Draw". Proses input skor ini akan berhenti secara otomatis jika pengguna memasukkan angka negatif, maka program akan berhenti. Kemudian, program akan melakukan perulangan kembali menggunakan fungsi len untuk menelusuri seluruh isi slice pemenang dan menampilin daftar hasilnya secara berurutan ke layar kemudian program akan print "Pertandingan selesai".

## Unguided 

### 4. 
Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk
membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa
apakah membentuk palindrom. Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi
untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.
*Palindrom adalah teks yang dibaca dari awal atau akhir adalah sama, contoh: KATAK, APA,
KASUR_RUSAK.

```go
package main

import (
    "fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
    var input string
    *n = 0
    for {
        fmt.Scan(&input)
        if input == "." {
            break
        }
        if *n < NMAX {
            (*t)[*n] = []rune(input)[0]
            *n++
        }
    }
}

func cetakArray(t tabel, n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("%c ", t[i])
    }
    fmt.Println()
}

func balikanArray(t *tabel, n int) {
    for i := 0; i < n/2; i++ {
        temp := (*t)[i]
        (*t)[i] = (*t)[n-1-i]
        (*t)[n-1-i] = temp
    }
}

func palindrom(t tabel, n int) bool {
    var tempTabel tabel = t
    balikanArray(&tempTabel, n)
    
    for i := 0; i < n; i++ {
        if t[i] != tempTabel[i] {
            return false
        }
    }
    return true
}

func main() {
    var tab tabel
    var m int

    fmt.Print("Teks         : ")
    isiArray(&tab, &m)

    var reversedTab tabel = tab
    balikanArray(&reversedTab, m)

    fmt.Print("Reverse teks : ")
    cetakArray(reversedTab, m)

    fmt.Printf("Palindrom    ? %t\n", palindrom(tab, m))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%204/output/soal2.png)
Jadi, program ini dibuat buat mengecek apakah sebuah kata kalau dibalik urutannya tetap sama atau disebut sebagai palindrom dengan menggunakan array bertipe rune maksimal 127 karakter. Cara kerja program ini yaitu pertama kamu diminta menginput teks karakter melalui fungsi isiArray yang akan terus berjalan sampai kamu mengetik tanda titik atau jumlahnya mencapai 127, di mana setiap karakter tersebut disimpan dalam indeks array yang urut. Setelah teks tersimpan, program mengeksekusi fungsi balikanArray yang fungsinya menukar posisi elemen dari depan ke belakang menggunakan variabel bantuan temp, lalu hasilnya ditampilkan ke layar lewat fungsi cetakArray. Terakhir, fungsi palindrom akan membandingkan susunan karakter asli dengan hasil yang sudah dibalik tersebut jika keduanya sama, maka program bakal mengembalikan nilai true.