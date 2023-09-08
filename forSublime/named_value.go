package main

import("fmt")

func getFullName() (firstName, middleName, lastName string){
	firstName = "William"
	middleName = "Bruce"
	lastName = "Bailey"

	return
}

func main(){
	axl, ro, se := getFullName()

	fmt.Println(axl, ro, se)
}