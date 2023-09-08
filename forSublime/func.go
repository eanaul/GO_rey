package main 

import("fmt")

func program(){
	var name string

	fmt.Println("oh halo, ini program percabangan menggunakan switch")
	fmt.Println("coba kamu masukin nama")
	fmt.Scan(&name)

	switch name {
	case "rehan":
		fmt.Println("halo admin, selamat datang")
	case "albert":
		fmt.Println("kamu bukan admin tapi boleh masuk")
	case "waluyo":
		fmt.Println("okey kamu member")
	case "thomas":
		fmt.Println("okey kamu member")
	default:
		fmt.Println("penyusup, get the hell out of here")
		return
	}


	var nis string

	fmt.Println("coba masukin nis kamu ya")
	fmt.Scan(&nis)

	switch length := len(nis); length == 8{
	case true: 
		fmt.Println("nis valid, silahkan")
	case false:
		fmt.Println("nis ga valid, yang bener we atuh")
	}
}

func main(){
	program()
}