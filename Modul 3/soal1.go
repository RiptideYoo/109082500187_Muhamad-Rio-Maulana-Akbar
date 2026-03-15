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