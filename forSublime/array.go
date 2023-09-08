package main
import "fmt"

func main(){
	city := [...]string{"Shibuya", "Minnesota", "Manchester", "Berlin", "Ho Chi Minh City"}

	fmt.Print(city, "\n")
	fmt.Print(city[4], "\n")

	city[4] = "Oregon"
	fmt.Println(city)

	fmt.Println("\n")

	// inisialisasi beberap array
	usa := [4]string{1:"New York", 3:"Indiana"}

	fmt.Println(usa, "\n")
	fmt.Println(len(usa))
}