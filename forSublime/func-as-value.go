package main

import("fmt")

func getGoodBye(name string) string {
	return "Selamat tinggal " + name
}

func main (){
	var x string

	fmt.Println("Tulis nama yang ingin kamu ucapkan selamat tinggal:")
	fmt.Scan(&x)

	sayGoodBye := getGoodBye
	result := sayGoodBye(x)
	fmt.Println(result)
}