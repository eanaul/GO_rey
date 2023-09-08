package main

import("fmt")

func whatName() (string, string, int){
	return "Rehan", "Aul", 12209113
}

func main(){
	firstName, lastName, yourNis := whatName()

	fmt.Println(firstName, lastName, "NIS:", yourNis)
}