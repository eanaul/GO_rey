// perlu diingat bahwa slice adalah potongan data yang ada di array

package main 

import("fmt")

func main(){
	var asean = [11]string{
		"Indonesia",
		"Malaysia", 
		"Singapure",
		"Thailand",
		"Myanmar", 
		"Laos", 
		"Vietnam", 
		"Philipines", 
		"Brunei", 
		"Kamboja", 
		"Timor Leste",
	}

	var slice1 = asean[4:8]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	asean[6] = "Sweden"
	fmt.Println(slice1)

	// program append slice

	days := [...]string{"Monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	daySlice1 := days[5:]
	daySlice1[0] = "new saturday"
	daySlice1[1] = "new sunday"

	fmt.Println(days) 

	daySlice2 := append(daySlice1, "today")
	daySlice2[0] = "kamari"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// kamu bisa buat slice dengan:

	newCountry := make([]string, 5, 5)

	newCountry[0] = "Iceland"
	newCountry[1] = "Norway"
	newCountry[2] = "Sweden"
	newCountry[3] = "Belgium"
	newCountry[4] = "Scotland"

	fmt.Println(newCountry)
	fmt.Println(len(newCountry))
	fmt.Println(cap(newCountry))
}