package main

import("fmt")

type Siswa struct{
	Nama, Rayon, Rombel string
	Nis 				int
}


func (siswa Siswa) sayCuy(nama string){
	fmt.Println("Halo", nama, "My name is", siswa.Nama)
}

func (a Siswa) sayHuuu(){
	fmt.Println("Huuuu from", a.Nama)
}


func main(){

	var ehan Siswa
	ehan.Nama = "Ehan Au"
	ehan.Nis = 12209113
	ehan.Rayon = "Cicurug 6"
	ehan.Rombel = "PPLG X-1"

	fmt.Println(ehan.Nama)
	fmt.Println(ehan.Nis)
	fmt.Println(ehan.Rayon)
	fmt.Println(ehan.Rombel)

	ehan.sayCuy("Joko")
	ehan.sayHuuu()

	// // struct literals

	// axl := Siswa{
	// 	Nama: "William Bruce Bailey",
	// 	Nis: 060262, 
	// 	Rayon: "Indiana, USA",
	// 	Rombel: "Guns N' Roses",
	// }

	// fmt.Println(axl)

	// // atau bisa seperti ini

	// roger := Siswa{"Roger Waters", "Bookham, UK", "Pink Floyd", 60942}

	// fmt.Println(roger)
}