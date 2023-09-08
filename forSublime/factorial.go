package main

import("fmt")

func cuy(nilai int) int{
	if nilai == 1{
		return 1
	}else{
		return nilai * cuy(nilai-1)
	}
}

func main(){
	fmt.Println(cuy(5))
}