package main

import "fmt"

func main() {
    var berat, d1, d2 int
    var total int

    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

    d1 = (berat/1000) 
    d2 = (berat % 1000)
    fmt.Println("detail berat: ", d1,"kg", " + ", d2,"gram")

    d1 = d1 * 10000

    if d2 >= 500{
        d2 = d2 * 5
    }else if d2 <= 500{
        d2 = d2 * 15
    }

    fmt.Println("detail biaya: ", "Rp. ", d1, " + ", "Rp. ", d2)
    total = d1 + d2
    fmt.Print("total biaya: ", total)

}
