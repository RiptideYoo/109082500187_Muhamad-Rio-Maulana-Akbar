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