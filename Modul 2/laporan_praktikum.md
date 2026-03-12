# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Muhamad Rio Maulana Akbar] - [109082500187]</p>

## Unguided 

### 1. 
Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main
import "fmt"
func main() {
	var (
	satu, dua, tiga string
	temp string
	)
	
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%202/output/soal1.png)
jadi pada pemprogaraman ini kita diminta inputan string sebanyak 3 kali, inputan pertama disimpan di variabel "satu", inputan kedua disimpan di variabel "dua" dan inputan ketiga disimpan di variabel "tiga". setelah inputan yang diminta sudah diisi maka program akan print output awal yaitu isi dari variabel satu, dua dan tiga (Output awal = satu, dua, tiga). setelah itu program akan melakukan penukaran isi variabel  (temp = satu, satu = dua, dua = tiga, dan tiga = temp) sehingga output akhir menghasilkan "Output akhir = dua tiga satu". 

## Unguided 

### 2. 
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.
#### soal2.go

```go
package main
import "fmt"

func main() {
	var s1, s2, s3, s4 string
	var berhasil bool

	berhasil = true

	for i:= 1; i<= 5; i++{
		fmt.Print("percobaan ", i, " :")
		fmt.Scan(&s1, &s2, &s3, &s4)

		if s1 != "merah" || s2 != "kuning" || s3 != "hijau"|| s4 != "ungu"{
			berhasil = false
		}

	}
	fmt.Print("BERHASIL : ", berhasil)
	
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%202/output/soal2.png)
Jadi cara kerja program di atas, pertama kita menyiapkan empat variabel string yaitu s1,
s2, s3, s4 untuk menampung warna dan satu variabel boolean berhasil yang awalnya
diset True. Kemudian kita masuk ke perulangan for yang akan berjalan sebanyak 5 kali
percobaan, di mana pada setiap putaran kita diminta menginputkan empat warna
sekaligus. Di dalam perulangan tersebut, ada kondisi if yang mengecek apakah urutan
warna yang diinputkan tepat sama dengan merah, kuning, hijau, dan ungu. Jika
pada salah satu percobaan urutannya salah, maka variabel berhasil akan berubah menjadi
false dan nilai ini akan tetap tersimpan meskipun pada percobaan berikutnya urutannya
benar. Setelah perulangan selesai melakukan 5 kali putaran, program akan berhenti dan
mencetak status akhir dari variabel tersebut, misalnya mencetak Berhasil :true jika
semua urutan warna tepat.

## Unguided 

### 3. 
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

```go
package main

import "fmt"

func main() {
    var berat, d1, d2 int
    var total int

    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

    d1 = (berat/1000) 
    d2 = (berat % 1000)
    fmt.Println("detail berat: ", d1,"kg", " + ", d2,"gram")

    d1 = d1 * 10000

    if d2 >= 500{
        d2 = d2 * 5
    }else if d2 <= 500{
        d2 = d2 * 15
    }

    fmt.Println("detail biaya: ", "Rp. ", d1, " + ", "Rp. ", d2)
    total = d1 + d2
    fmt.Print("total biaya: ", total)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%202/output/soal1.png)
Jadi cara kerja program kalkulator ini awalnya kita masukin inputan berat dalam sekala
gram contohnya 8500 atau setara 8.5 kg nah karena di soal kita diminta detail berat
yaitu berat utama dalam kg dan sisa berat dalam gram maka untuk berat utama kita bagi
seribu (berat/1000) dan sisa berat (berat%1000) kenapa modulus 1000?, karena untuk
memunculkan 3 digit belakang yaitu 500, terus berat dalam satuan kg kita kali dengan
10.000 karena biayanya 10.000 untuk 1kg kemudian untuk sisa berat kita menggunakan
pemisalan atau if, karena jika sisa berat kurang dari 500gram biayanya 15 sedangkan
jika lebih dari 500g biayanya cuma 5. Abis itu kita print harga berat utama dan sisa
berat terus kita print hasil penjumlahan harga berat utama dan sisa berat.
