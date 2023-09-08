package main

import("fmt")

func main(){
	for i := 1; i <= 10; i++{

		if i == 5{

			fmt.Println("gaada angkot no 5 ya dek")
			continue
		}

		fmt.Println("angkot", i)

	}
}