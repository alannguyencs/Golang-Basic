package main

import (
	"fmt"
	"time"	
)

func main() {
     message := make(chan string, 2) // buffer
     count := 5

     go func() {
          for i := 1; i <= count; i++ {
               fmt.Println("send message")
               message <- fmt.Sprintf("message %d", i)
          }
     }()

     

     for i := 1; i <= count; i++ {
		  time.Sleep(time.Second * 3)
          fmt.Println(<-message)
     }
}