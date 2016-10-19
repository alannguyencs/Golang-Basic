package main
import f"fmt"

type Human struct{
	name string
	age int
	weight int
}

type Student struct{
	Human //Student includes all fields of Human
	speciality string
}

func display(st Student){
	f.Println("Name: ", st.name);
	f.Println("Age: ", st.age);
	f.Println("Weight: ", st.weight);
	f.Println("Speciality: ", st.speciality);
	
}

func main(){
	mark := Student{Human{"Mark", 25, 64}, "Computer Science"}
	
	mark.age += 2
	mark.weight -= 1
	
	display(mark)
}