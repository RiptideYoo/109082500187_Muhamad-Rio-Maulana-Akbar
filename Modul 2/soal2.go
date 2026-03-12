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