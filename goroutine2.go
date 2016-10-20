package main

import "fmt"

func main() {
     // Start a goroutine and execute println concurrently
     go fmt.Println("goroutine message")
	
	for i := 0; i < 100; i++ {
		fmt.Println("main function message main function message main function message")
	}
     
}