package main
import "fmt"
func fx(x int)int {
	hasil:= x*x
	return hasil
}
func hx(x int)int{
	hasil := x + 1
	return hasil
}
func gx(x int)int{
	hasil := x-2
	return hasil
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
}