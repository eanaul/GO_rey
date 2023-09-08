package main

import("fmt")

func cuyAsoy(){
	fmt.Println("omomomomomo")
}

func runnApp(nilai bool){
	defer cuyAsoy()
	if nilai{
		panic("ERROR ALAMAK OIIII")
	}
	fmt.Println("Aplikasi berjalan, santuy")
}

func main(){
	runnApp(true)
}