package main

import "fmt"a

func main() {
     done := make(chan bool)

     go func() {
		 done <- true
          println("goroutine message")
		
          // We are only interested in the fact of sending itself, 
          // but not in data being sent.
          
     }()

     println("main function message")
	for i := 1; i < 1000; i++{
		fmt.Println(i)
	}
     <-done 
} 