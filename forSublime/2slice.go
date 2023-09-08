package main
import "fmt"

func main() {
	myCity := make([]string, 0)

	myCity = append(myCity, "California")
	myCity = append(myCity, "Philadelphia")
	myCity = append(myCity, "Oregon")

	fmt.Println(myCity)

	myBand := make([]string, 4, 4)

	myBand[0] = "Oasis"
	myBand[1] = "Led Zeppelin"
	myBand[2] = "Pink Floyd"
	myBand[3] = "The Beatles"

	fmt.Println(myBand)
}