package main
import "fmt"

func main(){
	arr := [10] int {1, 2, 3, 4, 5}
	var v, vv []int
	
	v = append(v, arr[0], arr[1], arr[4])
	fmt.Println(v[0])
	fmt.Println("length of slice v is ",len(v))
	fmt.Println("capacity of slice v is ",cap(v))
	
	vv = append(vv, 0)
	fmt.Println("length of slice vv is ",len(vv))
	fmt.Println("first element of vv: ", vv[0])
	copy(vv, v)                                    //copy elements of v to vv
	fmt.Println("length of slice vv is ",len(vv))
	fmt.Println("first element of vv: ", vv[0])
	
	fmt.Println(v)
}