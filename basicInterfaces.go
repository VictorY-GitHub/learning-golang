package main 

import (
	"fmt"
)

func myFunc (a interface{}) { // flexibility to pass in anything we want
	fmt.Println(a)
}

func main () {
	var age int = 25
	myFunc(age)
	var name string = "Victor Yun"
	myFunc(name)
}