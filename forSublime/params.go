package main

import("fmt")

func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func main(){
	sayHelloTo("Ehan", "Kopling")
	sayHelloTo("Udin", "Naget")
	sayHelloTo("Brook", "Lopez")
}