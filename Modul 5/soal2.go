package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	bintang(n, 1, 1)
}
func bintang(n, i, j int){
	if i > n{
		return 
	}
	if j <= i {
		fmt.Print("*")
		bintang(n, i, j+1)
	}else{
		fmt.Println()
		bintang(n, i+1, 1)
	}
}