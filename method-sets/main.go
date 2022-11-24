package main

import (
	"fmt"
	"strconv"
)

type szemely struct{
	Nev  string
	Kor  int	
	}

func (sz *szemely) beszel() string{
	return sz.Nev +" a nevem "+ strconv.Itoa(sz.Kor)+" eves vagyok es takaritok"
} 

func main() {
	magdaneni:=szemely{Nev : "Magda", Kor : 65}
	fmt.Println(magdaneni.beszel())
}


/*
● create a type person struct
● attach a method speak to type person using a pointer receiver
○ *person
● create a type human interface
○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
○ have it take in a human as a parameter
○ have it call the speak method
● show the following in your code
○ you CAN pass a value of type *person into saySomething
○ you CANNOT pass a value of type person into saySomething
● here is a hint if you need some help
*/