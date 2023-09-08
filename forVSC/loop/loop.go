package main

import("fmt")

func main(){
	// var x int

	// fmt.Println("kamu mau berapa kali nyetak angkot?")
	// fmt.Scan(&x)

	// for i := 1; i <= x; i++{
	// 	fmt.Println("ohayou, ini angkot ke-", i)
	// }

	slice := []string{"Reyhan", "Aulia", "Treeana", "William", "Bruce", "Bailey", "Josh", "Norton", "Steve", "Mark", "Axel", "James", "Parker", "Gord", "Derby", "Earnest", "Malcolm", "Paul"}

	// for i := 0; i < len(slice); i++{
	// 	fmt.Println(slice[i])
	// }

	for i, names := range slice{
		fmt.Println("Index", i, "Adalah", names)
	}

	for _, names := range slice{
		fmt.Println(names)
	}



}