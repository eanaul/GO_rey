package main

import("fmt")

type Blacklist func(string) bool

func periksaUsn(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("kamu diblok", name)
	}else{
		fmt.Println("silahkan", name)
	}
}

func main(){
	blacklist := func(name string) bool{
		return name == "rehan"
	}

	periksaUsn("Tobira", blacklist)
	periksaUsn("rehan", func(name string) bool{
		return name == "rehan"
	})
}