package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm(); //parse arguments, you have to call this by yourself
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	
	for k, v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	
	fmt.Fprintf(w, "Hello World!") //send data to client side 
}

func main(){
	http.HandleFunc("/", sayhelloName) //set router
	err := http.ListenAndServe(":9090", nil) //set listen socket
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
/*
After we execute the above code, the server begins listening to port 9090 in local host.
*/