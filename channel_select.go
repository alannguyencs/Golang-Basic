package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 1, 1
	fmt.Println("--------------------------------------------")
    for {
        select {
        		case c <- x:                 //set value of c == x
            		x, y = y, x + y
					//fmt.Println(<-c)
       		 	case <-quit:                //<-quit == true
        			fmt.Println("quit")
				//default:
					//fmt.Println("Default")
            	return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
		fmt.Println("********************************")
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}