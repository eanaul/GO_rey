package main

import "fmt"

func polaLingkaran(pola int) {
    var hasil string
    for i := 1; i <= pola + 1; i++ {
        for x := 1; x <= pola + 2; x++ {
            hasil += "🍕"
        }
        hasil += "\n"
    }
    fmt.Println(hasil)
}

func main() {

    var x int
    fmt.Println("Kamu mau lihat pola ke berapa?")
    fmt.Scan(&x)

    polaLingkaran(x)
}
