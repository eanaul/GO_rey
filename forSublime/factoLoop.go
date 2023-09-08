package main 

import("fmt")

func factorialLoop(nilai int) int{
	hasil := 1
	for i := nilai; i > 0; i--{
		hasil *= i
	}

	return hasil
}

func main(){
	loop := factorialLoop(7)

	fmt.Println(loop)
	fmt.Println(7 * 6 * 5 * 4 * 3 * 2 * 1)
}