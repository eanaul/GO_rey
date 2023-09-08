package main

import("fmt")

func coeg() interface{} {
	return false
}

func main(){
	result := coeg()
	// resultHasil := result.(string)
	// fmt.Println(resultHasil)

	// gunakan switch 
	switch value := result.(type){
	case string: 
		fmt.Println("Value", value, "is string" )
	case int:
		fmt.Println("Value", value, "is integer" )
	case bool:
		fmt.Println("Value", value, "is boolean" )
	default:
		fmt.Println("Unknown type" )
	}
}