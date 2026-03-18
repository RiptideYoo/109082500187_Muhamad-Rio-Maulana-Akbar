# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">[MUHAMAD RIO MAULANA AKBAR] - [109082500187]</p>

## Unguided 

### 1. 
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
#### soal1.go

```go
package main
import "fmt"
func faktorial(n int)int{
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++{
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int)int{
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int)int {
	return permutasi(n,r) / faktorial(r)
}
func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	if a >= c && b >= d {
		fmt.Printf("%d, %d\n", permutasi(a,c), kombinasi(a,c))
		fmt.Printf("%d, %d\n", permutasi(b,d), kombinasi(b,d))
	} else if c >= a && b>= d{
		fmt.Printf("%d, %d\n", permutasi(c,a), kombinasi(c,a))
		fmt.Printf("%d, %d\n", permutasi(b,d), kombinasi(b,d))
	}else if a >= c && d >= b {
		fmt.Printf("%d, %d\n", permutasi(a,c), kombinasi(a,c))
		fmt.Printf("%d, %d\n", permutasi(d,b), kombinasi(d,b))
	}else if c >= a && d>= b{
		fmt.Printf("%d, %d\n", permutasi(c,a), kombinasi(c,a))
		fmt.Printf("%d, %d\n", permutasi(d,b), kombinasi(d,b))
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%203/output/soal1.png)
Jadi, program diatas merupakan program yang berfungsi sebagai kalkulator untuk menghitung nilai permutasi dan kombinasi dari 4 angka yang kita inputkan, nilai inputan tadi akan diproses menggunakan fungsi faktorial untuk perhitungannya. Di dalam fungsi main terdapat if-else untuk membaca angka mana yang lebihbesar pada tiap pasangan agar perhitungannya tidak salah dimana nilai n harus lebih besar atau sama dengan r, setelah sudah ditentukan pada if-else maka program akan mengeluarkan output nilai hasil permutasi dan kombinasi.

## Unguided 

### 2. 
Diberikan tiga buah fungsi matematika yaitu f (x) = x^2
, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x)
dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),
dan baris ketiga adalah (hofog)(c)!
#### soal2.go

```go
package main
import "fmt"
func fx(x int)int {
	hasil:= x*x
	return hasil
}
func hx(x int)int{
	hasil := x + 1
	return hasil
}
func gx(x int)int{
	hasil := x-2
	return hasil
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%203/output/soal1.png)
Jadi, Program ini dibuat untuk menjalankan fungsi komposisi. Kita diminta 3 inputan yang dimasukkan ke variabel a, b, c yang kemudian diproses dengan menggunakan 3 fungsi berbeda, Fx untuk menguadratkan angka, hx untuk menambahkan angka dengan satu dan gx untuk mengurangkan angka dengan 2.Urutan eksekusinya dimulai dari fungsi paling dalam kurung menuju ke luar contohnya fogoh(x) pada fungsi ini variabel a ditambah 1, dikurang 2, lalu hasilnya dikuadratkan.

## Unguided 

### 3. 
[Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius
r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)
func jarak(a, b, c, d float64)float64{
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	}else if dalam2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%203/output/soal3.png)
Jadi, program ini dibuat untuk mendeteksi posisi titik di antara dua buah lingkaran yang berbeda dengan menggunakan rumus jarak matematika.Cara kerjanya pertama kita inputkan koordinat pusat dan jari-jarinya untuk dua lingkaran dan juga koordinat satu titik sembarang, nah dari fungsi jarak program ngitung berapa jauh titik tersebut dari titik pusatnya masing-masing, abis itu fungsi didalam bakal nentuin jaraknya lebih kecil atau sama dengan jari-jari lingkaran. 