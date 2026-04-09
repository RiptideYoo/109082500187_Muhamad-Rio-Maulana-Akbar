# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>
<p align="center"> Muhamad Rio Maulana Akbar - 109082500187 </p>

## Unguided 

### 1. 
Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai
suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat
diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku
ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci
tersebut.
#### soal1.go

```go
package main
import "fmt"
func fibbonaci(n int)int{
	if n == 0{
		return 0
	}else if n==1 {
		return 1
	}
	return fibbonaci(n-1) + fibbonaci(n-2)
}
func main() {
	var n int

	fmt.Scan(&n)

	for i:= 0; i <= n; i++{
		fmt.Print(fibbonaci(i)," ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
Jadi program ini dibuat untuk nampilin deret angka Fibonacci mulai dari suku ke-0 sampai ke-n sesuai dengan angka yang kita input, di mana proses perhitungannya pakai fungsi rekursif bernama fibbonaci. Di dalam fungsi itu, ada If-else yang nentuin kalau inputnya 0 bakal balik ke 0 dan kalau 1 bakal balik ke 1, sedangkan buat angka di atas itu, program bakal manggil dirinya sendiri berkali-kali buat ngejumlahin dua angka sebelumnya (n-1 ditambah n-2). selanjutnya, di fungsi main, program bakal nge-loop dari angka 0 sampai batas n yang lo masukin tadi, lalu nyetak hasil perhitungan tiap sukunya ke samping biar ngebentuk deret yang rapi.


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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
Jadi program ini dibuat untuk mencari faktor dari angka. pertama kita diminta inputan yang nantinya disimpan ke variabel n, setelah itu program bakal memanggil func faktor. Didalam func faktor terdapat if untuk pengecekan apakah nilai i lebih dari n jika iya maka program akan berhenti, selanjutnya masuk ke if yang kedua dimana nilai n akan di modulus oleh nilai i, jika sisanya 0 maka itulah faktor dari n.
