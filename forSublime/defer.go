package main 

import("fmt")

func lumber(){
	fmt.Println("function berhasil dijalankan!")
}

func runApp(nilai int){
	defer lumber()
	fmt.Println("application running")
	result := 100 / nilai
	fmt.Println(result)
}

func main(){
	runApp(5)
}