package main

import("fmt")

func main(){

	siswa := map[string] string{
		"nama": "ehan",
		"nis": "12209113",
		"rombel": "pplg 1",
		"rayon": "cicurug 6",
		"salah": "oops",
		"title": "Pecinta sholawat",
	}
	
	fmt.Println(siswa)
	fmt.Println(len(siswa))

	siswa["title"] = "jago koding"

	delete(siswa, "salah")

	fmt.Println(siswa)
	fmt.Println(len(siswa))
}