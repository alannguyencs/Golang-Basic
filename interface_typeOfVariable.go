package main

import(
	"fmt"
	"strconv"
)

type Element interface{}   //interface for everything
type List []Element

type Person struct {
	name string
	age int
}

func (p Person) String() string{
	return "(name: " + p.name + " - age:     " + strconv.Itoa(p.age) + " years)"
}

func main(){
	list := make(List, 3) //initialize vector List
	
	list[0] = 1  //an int
	list[1] = "Hello" //a string
	list[2] = Person{"Alan", 25}
	
	for index, element := range list{
		if _, is_int := element.(int); is_int{ //check if it is int
			fmt.Printf("%d is Integer \n", index)
		}
		
		if _, is_string := element.(string); is_string{
			fmt.Printf("%d is String \n", index)
		}
		
		if value, is_person := element.(Person); is_person{
			fmt.Println(value)
		}
	}
	
	//switch case
	
	for _, element := range list{
		switch value := element.(type){
			case int:
				fmt.Println("Integer")
			case string:
				fmt.Println("String")
			case Person:
				fmt.Println(value)
			default:
				fmt.Println("Other type")
		}
	}
}