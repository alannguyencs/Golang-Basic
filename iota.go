package main
import "fmt"

const(
	x = iota //x = 0
	y = iota //y = 1
	z = iota //z = 2
	w //if there is no expression after the constants name, it uses the last expression (w = iota)
)

const v = iota //once iota meets keyword `const`, it resets to 0, then v = 0

const(
	e, f, g = iota, iota, iota // e = f = g = 0 values of iota are the same in one line	
)

func main(){
	fmt.Print("w = ", w);
}