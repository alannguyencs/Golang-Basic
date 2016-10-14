package main
import "fmt"

func main(){
	a := [3] int {1, 2, 4}
	b := [...] int {4, 5, 7} //use ... to replace the length parameter and go will calculate it for you
	c := [10] int {1, 2, 4} //the first 3 elements have been declared
	
	//2-dimensional array
	easyArray := [2][4] int {{1, 2, 3, 4}, {5, 6, 7, 8}}
	
	
	fmt.Println(a[0])
	fmt.Println(b[1])
	fmt.Println(c[2])
	fmt.Println(easyArray[0][0])
}