package main

import("fmt")

func Ups(i int) interface{} {
	if i == 1{
		return 1
	} else if i == 2{
		return true 
	} else{
		return "Eh? ups.."
	}
}

func main() {
	data := Ups(4)
	fmt.Println(data)
}