package main

import("fmt")

func main(){
	var nama string

	fmt.Println("masukan nama dulu yo")
	fmt.Scan(&nama)

	if name := nama; name == "rehan"{
		fmt.Println("kamu admin, selamat datang")
	} else if name == "suryo"{
		fmt.Println("kamu bukan admin, tapi boleh masuk, selamat datang")
	} else{
		fmt.Println("penyusup, keluar sebelum aku panggil police")
	}
}