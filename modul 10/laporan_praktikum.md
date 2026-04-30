# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center"> Muhamad Rio Maulana Akbar - 109082500187 </p>

## Unguided 

### 1. 
Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar.
Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak
kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan
bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya
N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.
Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan
terbesar.
#### soal1.go

```go
package main
import "fmt"
func main() {
	var n int
	var kelinci [1000]float64
	var min float64
	var max float64

	fmt.Scan(&n)
	if n > 1000{
		n = 1000
	}

	for i:= 0; i<n; i++{
		fmt.Printf("masukkan berat kelinci ke %d: ", i+1)
		fmt.Scan(&kelinci[i])
	}

	min = kelinci[0]
	max = kelinci[0]

	for i:= 1;i < n ;i++{
		if kelinci[i] < min{
			min = kelinci[i]
		}
		if kelinci[i] > max{
			max = kelinci[i]
		}
	}
	fmt.Printf("berat kelinci terkecil adalah %.2f dan berat terbesar adalah %.2f",min,max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%2010/output/soal1.png)
jadi, program diatas dibuat untuk mencari berat terkecil dan terbesar dari kelinci menggunakan array yang isinya sampe 1000 (geloo), cara kerjanya pertama kita diminta inputan banyaknya keclinci yang mau dicompare beratnya misal nih aku input 5 nanti program bakal minta berat kelinci sebanyak 5, setelah kita input masing-masing beratnya nanti program bakal masuk ke perulangan dari 1 sampai ke n, jika isi array kelinci[i] lebih kecil dari min maka nilai min = kelinci[i] begitu juga dengan nilai max. setelah mendapat nilai min dan max peogram bakal print output berat kelinci paling kecil dan terbesar sekian terimakasih ALLAHUAKBAR. 


## Unguided 

### 2. 
Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program
ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan
dijual.
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan
y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya
ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil
yang menyatakan banyaknya ikan yang akan dijual.
Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan
total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang
dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah
sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### soal2.go

```go
package main
import "fmt"
func main() {
	var x, y int
	var BeratIkan [1000]float64
	var totalBeratWadah float64
	var jumlahIkan int
	var wadah []float64
	var totalBeratsemuaWadah float64
	fmt.Print("masukkan jumlah ikan yang mau dijual dan dimasukkan ke wadah: ")
	fmt.Scan(&x, &y)

	for i:= 0; i < x; i++{
		fmt.Printf("masukkan berat ikan %d : ", i+1)
		fmt.Scan(&BeratIkan[i])
	}
	for i:= 0; i < x;i++{
		totalBeratWadah += BeratIkan[i]
		jumlahIkan ++

		if jumlahIkan == y || i == x-1{
			wadah = append(wadah, totalBeratWadah)
			totalBeratWadah = 0
			jumlahIkan = 0 
		}
	}
	for i:= 0; i < len(wadah); i++{
		fmt.Printf("berat wadah ke %d : %.2f\n", i+1, wadah[i])
		totalBeratsemuaWadah += wadah[i]
	}
	rata_rata:= totalBeratsemuaWadah / float64(len(wadah))
	fmt.Printf("rata-ratanya: %f\n", rata_rata)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%2010/output/soal2.png)
Jadi, program di atas dibuat untuk menghitung total berat ikan per wadah dan rata-ratanya menggunakan array yang ukurannya sampai 1000 (geloo), cara kerjanya pertama kita diminta input dua nilai, yaitu jumlah total ikan dan kapasitas maksimal satu wadah. Misal nih aku input 5 dan 2, berarti ada 5 ikan dan tiap wadah cuma bisa nampung 2 ikan. Setelah itu, program bakal minta kita input berat masing-masing ikan sebanyak 5 kali. Nah, pas masuk ke proses perulangan, program bakal menjumlahkan berat ikan satu-satu ke dalam wadah. Kalau jumlah ikan di wadah sudah mencapai batas (tadi kita input 2) atau ikannya sudah habis, maka total berat di wadah itu bakal disimpan ke dalam slice aka array kosong bernama wadah (asekk), terus hitungan beratnya di-reset lagi jadi nol buat wadah berikutnya. setelah semua wadah terisi, program bakal nge-print berat masing-masing wadah secara berurutan dan menghitung rata-rata berat dari semua wadah yang ada, terus hasilnya langsung dimunculin di layar. Sekian terimakasih ALLAHUAKBAR TAKBIR.

## Unguided 

### 3. 
Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data
berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data
yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var dataBerat arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	if n > 100 {
		n = 100
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	hitungMinMax(dataBerat, n, &min, &max)

	hasilRerata := rerata(dataBerat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", hasilRerata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%2010/output/soal3.png)
Jadi, program ini intinya buat nyari berat balita paling kecil, paling gede, sama rata-ratanya pakai array yang gedenya sampai 100 (edun!), cara kerjanya pertama kita diminta input dulu mau masukin berapa data balita, misal nih aku input 5 nanti program bakal minta kita ketik berat balita sebanyak 5 kali terus disimpen di array, nah abis itu program bakal manggil fungsi hitungMinMax buat ngebandingin satu-satu isinya kalau ada yang lebih kecil dari min atau lebih gede dari max bakal langsung ditandai, terus ada juga fungsi rerata buat ngejumlahin semua beratnya terus dibagi sama jumlah balitanya, dan terakhir program bakal langsung print output berat paling kecil, paling gede, sama rata-ratanya sekian terimakasih ALLAHUAKBAR.
