package main

import("fmt")

func main(){

	// do{
		var(
			a int
			b int
			op int
			yn string
		)

		fmt.Println("masukan angka pertama: ")
		fmt.Scan(&a)
		fmt.Println("masukan angka kedua: ")
		fmt.Scan(&b)

		fmt.Println("pilih operasi: ")
		fmt.Println("1 / + , 2 / - , 3 : /, 4 / *")
		fmt.Scan(&op)

		if op == 1{
			fmt.Println(a + b)
		}else if op == 2{
			fmt.Println(a - b)
		}else if op == 3{
			fmt.Println(a / b)
		}else if op == 4{
			fmt.Println(a * b)
		}else{
			fmt.Println("gausa ngaco")
		}

		fmt.Println("mau mengulang? y/n")
		fmt.Scan(&yn)


	// }while yn == y


}