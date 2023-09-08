package main

import("fmt")

func main(){

	

// fmt.Println("berapa kali?")
// fmt.Scan(&i)

// if i == 1{
// 	fmt.Println("kamu memasukan angka 1")
// } else if i == 2{
// 	fmt.Println("kamu memasukan angka 2")
// } else if i == 3{
// 	fmt.Println("kamu memasukan angka 3")
// }else{
// 	fmt.Println("kamu bukan memasukan angka 1-3")
// }

for i := 0; i <= 10; i++{
	fmt.Println(i)
}

 var b int

fmt.Println("masukan angka dari 0-5")
fmt.Scan(&b)

city := [5]string{"Jakarta", "Shibuya", "Berlin", "Memphis", "Stockholm"}

fmt.Println(city[b])
}