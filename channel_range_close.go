package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
	fmt.Println("go")
    x, y := 1, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x + y
		fmt.Println(i, " go")
    }
    close(c)
}

func main() {
    c := make(chan int, 10)
	
	fmt.Println("capacity of c: ", cap(c))
    go fibonacci(cap(c), c)
	
	fmt.Println("main")
	
    for i := range c {   //will not stop reading data from channel until the channel is closed
        fmt.Println(i)
    }
}