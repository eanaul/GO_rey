package main

import("fmt")

type Foulter func(string) string

func filter(name string, sensor Foulter){
	sudahFilter := sensor(name)
	fmt.Println("Hallo,", sudahFilter)
}

func diFilter(name string) string{
	if name == "rehan"{
		return "Bosen han sumpah"
	} else{
		return name
	}
}

func main(){
	var x string
	fmt.Scan(&x)

	filter(x, diFilter)
}