package main
import "fmt"

func main(){
	digit := make(map[string] int)
	
	digit["one"] = 1
	digit["five"] = 5
	digit["eight"] = 8
	
	fmt.Println("The eighth number is: ", digit["eight"])
	fmt.Println("Length of this map: ", len(digit))
	
	delete(digit, "eight")
	fmt.Println("After deleted, the eighth number is: ", digit["eight"])
	
	
	/*
	map is a reference type. If two map point to same underlying data, any change will affect both of them
	*/
	
	m := make(map[string] string)
	m["Hello"] = "Bonjour"
	m1 := m               //2 map -- same underlying data
	m1["Hello"] = "Salut"
	fmt.Println(m["Hello"]) //---> Salut
}