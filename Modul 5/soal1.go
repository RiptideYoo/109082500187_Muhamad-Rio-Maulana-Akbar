package main
import "fmt"
func fibbonaci(n int)int{
	if n == 0{
		return 0
	}else if n==1 {
		return 1
	}
	return fibbonaci(n-1) + fibbonaci(n-2)
}
func main() {
	var n int

	fmt.Scan(&n)

	for i:= 0; i <= n; i++{
		fmt.Print(fibbonaci(i)," ")
	}
}