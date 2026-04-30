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
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%210/output/soal1.png)
jadi, program diatas dibuat untuk mencari berat terkecil dan terbesar dari kelinci menggunakan array yang isinya sampe 1000 (geloo), cara kerjanya pertama kita diminta inputan banyaknya keclinci yang mau dicompare beratnya misal nih aku input 5 nanti program bakal minta berat kelinci sebanyak 5, setelah kita input masing-masing beratnya nanti program bakal masuk ke perulangan dari 1 sampai ke n, jika isi array kelinci[i] lebih kecil dari min maka nilai min = kelinci[i] begitu juga dengan nilai max. setelah mendapat nilai min dan max peogram bakal print output berat kelinci paling kecil dan terbesar sekian terimakasih ALLAHUAKBAR. 


## Unguided 

### 2. 
Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan
menggunakan fungsi rekursif. N adalah masukan dari user.
#### soal2.go

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	bintang(n, 1, 1)
}
func bintang(n, i, j int){
	if i > n{
		return 
	}
	if j <= i {
		fmt.Print("*")
		bintang(n, i, j+1)
	}else{
		fmt.Println()
		bintang(n, i+1, 1)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%205/output/soal2.png)
Jadii program ini dibuat untuk nunjukin bintang bentuk segitiga siku. pertama kita diminta input angka yang bakal disimpan ke variabel n, setelah itu program bakal jalanin atau manggil func bintang. didalam func bintang terdapat if-else, if yang pertama berfungsi untuk mengecek apakah nilai i itu lebih besar dari n atau engga kalau kondisi i belum melebihi n, program bakal masuk ke pengecekan kedua untuk nentuin apakah harus nyetak bintang atau ganti baris, di mana selama nilai j masih kurang dari atau sama dengan i, program bakal nyetak karakter bintang ke samping dan manggil dirinya sendiri lagi sambil nambahin nilai j (+1). Tapi, kalau si j sudah lebih besar dari program bakal jalanin fmt.Println() buat pindah ke baris baru dan manggil kembali fungsi bintang dengan nambahin nilai i (+1) serta ngereset nilai j balik ke 1. Proses ini bakal ngulang sampai jumlah barisnya pas sesuai angka n yang kita masukin tadi, sehingga terbentuklah pola segitiga siku-siku dari bintang-bintang tersebut. (maaf yapping :v)

## Unguided 

### 3. 
Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari
suatu N, atau bilangan yang apa saja yang habis membagi N.
#### soal3.go

```go
package main
import "fmt"
func faktor(n, i int){
	if i>n{
		return
	}
	if n%i == 0{
		fmt.Print(i, " ")
		
	}
	faktor(n, i+1)
}
func main() {
	var n int
	fmt.Scan(&n)
	faktor(n,1)
	
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%205/output/soal3.png)
Jadi program ini dibuat untuk mencari faktor dari angka. pertama kita diminta inputan yang nantinya disimpan ke variabel n, setelah itu program bakal memanggil func faktor. Didalam func faktor terdapat if untuk pengecekan apakah nilai i lebih dari n jika iya maka program akan berhenti, selanjutnya masuk ke if yang kedua dimana nilai n akan di modulus oleh nilai i, jika sisanya 0 maka itulah faktor dari n.
