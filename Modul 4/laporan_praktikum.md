# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">Muhamad Rio Maulana Akbar - 109082500187</p>

## Unguided 

### 1. 
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
#### soal1.go

```go
package main
import "fmt"
func factorial(n int, hasil *int){
	*hasil = 1
	for i := 1; i <= n; i++{
		*hasil = *hasil * i
	}

}
func permutation(n, r int, hasil *int ){
	var faktorialAtas, faktorialBawah int
		factorial(n, &faktorialAtas)
		factorial(n-r, &faktorialBawah)

		*hasil = faktorialAtas / faktorialBawah 
	
}
func combination(n, r int, hasil *int){
	var faktorialAtas, faktorialBawah, rFaktorial int

		factorial(n, &faktorialAtas)
		factorial(r, &rFaktorial)
		factorial(n-r, &faktorialBawah)

		*hasil = faktorialAtas / (rFaktorial*faktorialBawah)
	
}

func main() {
	var a, b, c, d int
	var a1, a2, b1, b2 int
	fmt.Print("masukkan 4 buah angka: ")
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &a1)
	combination(a, c, &a2)

	permutation(b, d, &b1)
	combination(b, d, &b2)

	fmt.Println(a1, a2)
	fmt.Print(b1, b2)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
Jadi program ini dibuat untuk mencari hasil permutasi dan kombinasi. pertama kita diminta untuk input 4 angka yaitu a, b, c, d yang dimana nantinya kita gunakan untuk mencari permutasi dari a dan c, permutasi b dan d begitupun untuk kombinasi. setelah memasukkan input program akan menjalankan func permutation yang dimana sebelum masuk ke func permutation akan diproses terlebih dahulu ke func faktorial untuk mencari nilai faktorial dari nilai masing-masing variabel (a,b,c,d). Setelah itu program akan lanjut ke func permutation untuk mencari nilai permutasinya, hasil permutasi ini akan disimpan ke variabel b1 yang nantinya akan di print untuk menunjukan nilai permutasinya, proses untuk kombinasi juga tahapannya sama seperti permutasi.

## Unguided 

### 2. 
Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal
yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan
soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program
harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total
soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.
Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca
di dalam prosedur.
#### soal1.go

```go
package main
import "fmt"
func factorial(n int, hasil *int){
	*hasil = 1
	for i := 1; i <= n; i++{
		*hasil = *hasil * i
	}

}
func permutation(n, r int, hasil *int ){
	var faktorialAtas, faktorialBawah int
		factorial(n, &faktorialAtas)
		factorial(n-r, &faktorialBawah)

		*hasil = faktorialAtas / faktorialBawah 
	
}
func combination(n, r int, hasil *int){
	var faktorialAtas, faktorialBawah, rFaktorial int

		factorial(n, &faktorialAtas)
		factorial(r, &rFaktorial)
		factorial(n-r, &faktorialBawah)

		*hasil = faktorialAtas / (rFaktorial*faktorialBawah)
	
}

func main() {
	var a, b, c, d int
	var a1, a2, b1, b2 int
	fmt.Print("masukkan 4 buah angka: ")
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &a1)
	combination(a, c, &a2)

	permutation(b, d, &b1)
	combination(b, d, &b2)

	fmt.Println(a1, a2)
	fmt.Print(b1, b2)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
Jadi program ini dibuat untuk mencari siapa pemenang perlombaan coding. pertama kita diminta input nama peserta lomba dan dan waktu pengerjaan soal sebanyak 8 kali. Program ini bakal terus minta input nama peserta satu per satu sampai kita ketik "Selesai" untuk mengakhiri loop, di mana setiap nama yang masuk bakal langsung diproses lewat subprogram hitungSkor buat ngitung berapa soal yang berhasil dijawab (maksimal 8 soal) dengan syarat waktu pengerjaannya gak boleh lebih dari 300 menit. Di dalam fungsi tersebut, kita pakai pointer supaya jumlah soal dan total skornya langsung terupdate ke variabel utama, lalu program bakal ngecek: kalau peserta itu punya jumlah soal lebih banyak atau soalnya sama tapi total waktunya lebih kecil (lebih cepat), maka nama dia bakal disimpan sebagai pemenang sementara. Setelah semua data diinput, program bakal nampilin nama pemenang beserta rekor jumlah soal dan total skor yang didapat.