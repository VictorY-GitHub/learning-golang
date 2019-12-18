package main

import (
	"fmt"
)

func main () {
	var myInt int8
	for i := 0; i < 129; i++ {
		myInt += 1	
	}
	fmt.Println(myInt) // overflows to -127, usually just use "int"

	var a uint8 = 5
	var b int16 = 6
	fmt.Println(int(a)+int(b)) // casting in Go

	// var f float32
	// var f2 float64 // floating point variables

	var flag bool = true 
	var flag2 bool = false // bool variables
	if flag && flag2 { // selection structures
		fmt.Println("1 && 0 = false") // won't run
	}
	if flag || flag2 {
		fmt.Println("1 || 0 = true") // will run
	}

	var s string = "Victor Yun"
	for i:=0; i < len(s); i++ {
		fmt.Println(s[i]) // prints out ASCII values of the characters (with \n)
	}

	for i:=0; i < len(s); i++ {
		fmt.Println(string(s[i])) // need to cast strings in order to print char
	}

}