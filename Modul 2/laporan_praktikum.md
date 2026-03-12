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
![Screenshot Output Unguided 1_1](https://github.com/RiptideYoo/109082500187_Muhamad-Rio-Maulana-Akbar/blob/main/Modul%202/output/soal1.png)

