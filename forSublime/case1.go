package main

import("fmt")


func main(){

	// slice kosong
	var city []string

	fmt.Println("Masukan Nama Kota: ")

	for i:= 0 ; i < 5; i++ {
		var cit string
		_, err := fmt.Scan(&cit)

		if err != nil {
			break
		}

		city = append(city, cit)
	}

	fmt.Println("kota yang dimasukan: ")
	fmt.Println(city)

}