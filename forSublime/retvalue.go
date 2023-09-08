package main

import("fmt")

func getHello(name string) string{
	return "halooo " + name
}

func main(){
	hasil := getHello("ehaan")
	fmt.Println(hasil)
}