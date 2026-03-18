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
[penjelasan]

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
[penjelasan]