package main

import "fmt"
/*
	ch := make(chan type, n)

	n == 0 ! non-buffer（block）
	n > 0 ! buffer（non-block until n elements in the channel）
*/
func main() {
    c := make(chan int, 2)  // change 2 to 1 will have runtime error, but 3 is fine
	/*seem to be a slice of channels*/
    c <- 1
    c <- 2
    fmt.Println(<-c)
    fmt.Println(<-c)
}