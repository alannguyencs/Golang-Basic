package main
import "fmt"

func add(a *int, k int) int{
	*a = *a + k  //we change value of a
	return *a    //return new value of a
}

func main(){
	x := 3
	y := add(&x, 5)
	
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}