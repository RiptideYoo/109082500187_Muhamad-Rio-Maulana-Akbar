package main
import "fmt"
type pusat struct{
	x, y int
}
type lingkaran struct{
	titik pusat
	r int
}
func main() {
	var l1, l2 lingkaran
	var d pusat
	var Circle1, Circle2 bool

	fmt.Scanln(&l1.titik.x, &l1.titik.y, &l1.r)
	fmt.Scanln(&l2.titik.x, &l2.titik.y, &l2.r)
	fmt.Scan(&d.x, &d.y)

	Circle1 = (d.x - l1.titik.x)*(d.x - l1.titik.x) + (d.y - l1.titik.y)*(d.y - l1.titik.y) < l1.r*l2.r
	Circle2 = (d.x - l2.titik.x)*(d.x - l2.titik.x) + (d.y - l2.titik.y)*(d.y - l2.titik.y) < l2.r*l1.r

	if Circle1 && Circle2 {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	}else if Circle1{
		fmt.Print("Titik di dalam lingkaran 1")
	}else if Circle2{
		fmt.Print("Titik di dalam lingkaran 2")
	}else{
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}



}