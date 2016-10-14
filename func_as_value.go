package main
import "fmt"

type testInt func(int) bool //define a function type of variable

func isOdd (x int) bool {return (x % 2 == 1)}
func isEven (x int) bool {return (x % 2 == 0)}

func filter(v [] int, fn testInt) []int{
	var ans []int
	for _, value := range(v){
		if(fn(value)) {ans = append(ans, value)}
	}
	return ans
}

func main(){
	v := []int {1, 1, 2, 3, 5, 8}
	fmt.Println("v = ", v)
	
	odd := filter(v, isOdd)
	fmt.Println("Odd = ", odd)
	
	even := filter(v, isEven)
	fmt.Println("Even = ", even)
}

