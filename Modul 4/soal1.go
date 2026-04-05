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