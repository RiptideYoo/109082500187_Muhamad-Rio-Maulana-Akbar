package main
import "fmt"
func main() {
	var x, y int
	var BeratIkan [1000]float64
	var 
	fmt.Print("masukkan jumlah ikan yang mau dijual dan dimasukkan ke wadah: ")
	fmt.Scan(&x, &y)

	for i:= 0; i < x; i++{
		fmt.Printf("masukkan berat ikan %d", i+1)
		fmt.Scan(&BeratIkan[i])
	}


}