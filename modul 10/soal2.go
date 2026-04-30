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