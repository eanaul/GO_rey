package main

import "fmt"

func main(){
	x:= 50

	if x == 50{
		fmt.Println("Germany")
	} else if x == 100{
		fmt.Println("France")
	} else{
		fmt.Println("Belgium")
	}


	// menghitung jumlah huruf

	name := "Ean"
	e := name[2]
	eString := string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}