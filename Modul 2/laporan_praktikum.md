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
