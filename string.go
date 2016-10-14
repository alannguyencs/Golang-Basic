package main
import "fmt"

func main(){
	s := "hello"  //can not change value of s[0-4]
	c := []byte(s) //convert s to character array
	c[0] = 'X'       //change value of c[0]
	s2 := string(c) //convert character array to string
	fmt.Println(s2)
	
	s = "X" + s[1:] //get values of s
	fmt.Println(s) 	
	//-----------------------------------------------------------
	t := `hello
	world`               //it gets multiline string
	fmt.Println(t)
}
