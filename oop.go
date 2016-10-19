package main
import f"fmt"

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YEALLOW
)

type Color byte    //special struct


type Box struct{
	width, height, depth float64
	color Color
}

type BoxList []Box  //a slice of boxes

func (b *Box) Volume() float64{  //write OOP
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color){
	b.color = c
}

func (bl BoxList) BiggestsColor() Color{  //find color of the biggest Box     BiggestsColor is a method of BoxList 
	mx := -1.0
	k := Color(1) // k = 1
	for _, b := range bl{
		if b.Volume() > mx{
			mx = b.Volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack(){
	for i,_ := range bl{
		bl[i].SetColor(BLACK)
	}
	//for _, b := range bl{     //----> b is a copy of new box, can not be used to change object
	//	b.SetColor(BLACK)
	//}
}
 
func (c Color) String() string{ //convert number to String    ------> String is a method of Color
	strings := []string {"White", "Black", "Blue", "Red", "Yeallow"}
	return strings[c]
}

func main(){
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{5, 20, 15, WHITE},
		Box{4, 6, 4, BLACK},
		Box{4, 4, 2, BLUE},
	}
	
	f.Printf("We have %d boxes in our set \n", len(boxes))
	
	f.Println("The volume of the first box: ", boxes[0].Volume())
	
	f.Println("The color of the last box: ", boxes[len(boxes) - 1].color.String())
	
	f.Println("The biggest one is: ", boxes.BiggestsColor().String())
	
	boxes.PaintItBlack();
	f.Println("Obviously, now the biggest one is: ", boxes.BiggestsColor().String())
	
}