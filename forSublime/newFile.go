package main

import("fmt")

func sayHi(name string){
	fmt.Println("Hi ", name)
}

func main(){
	var x string

	fmt.Scan(&x)
	sayHi(x)
}