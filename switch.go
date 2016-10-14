package main
import "fmt"

func main(){
	x := 6
	switch x{
		case 1: 
			fmt.Println("haaaaaa")
			fallthrough                   //if you want to match more cases
		case 6:
			fmt.Println("6666666666")
			fallthrough
		case 8:
			fmt.Println("88888888888")
			fallthrough
		default:
		fmt.Println("default case")
	}
	
	switch x{
		case 1: 
			fmt.Println("haaaaaa")
			//fallthrough                   //if you want to match more cases
		case 6:
			fmt.Println("6666666666")
			//fallthrough
		case 8:
			fmt.Println("88888888888")
			//fallthrough
		default:
		fmt.Println("default case")
	}
}