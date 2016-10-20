package main

import "fmt"

func main(){
	 message := make(chan string)
     count := 3

     go func() {
          for i := 1; i <= count; i++ {
               message <- fmt.Sprintf("message %d", i)
          }
		 close(message)
     }()

     for msg := range message {   //this will run until message is closed.
          fmt.Println(msg)
     }
}
/*
make 1 parallel processes:
send message1 to channel
							receive message1
send message2 to channel
							receive message2
send message3 to channel
							receive message3
close channel
							finish loop
*/