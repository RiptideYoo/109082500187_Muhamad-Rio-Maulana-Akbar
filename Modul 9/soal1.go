package main
import (
	"fmt"
	"math"
)
type titik struct{
	x, y int
}
type lingkaran struct{
	pusat titik
	r int
}
func jarak(p, q titik)float64{
	x := (p.x - q.x)*(p.x - q.x)
	y := (p.y - q.y)*(p.y - q.y)
	return math.Sqrt(float64(x)+float64(y))
}
func didalam(c lingkaran, p titik)bool{
	return jarak(c.pusat, p) < float64(c.r)
}
func main() {
	var l1, l2 lingkaran
	var d titik
	var Circle1, Circle2 bool

	fmt.Scanln(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scanln(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&d.x, &d.y)

	Circle1 = didalam(l1, d)
	Circle2 = didalam(l2, d)

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