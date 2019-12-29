package main

import (
	"fmt"
)

func myfunction(firstName string, lastName string) (string) { // parameters, results
	fullName := firstName + " " + lastName // concatenation
	return fullName
}

// it is also possible to have 2 results returned from a function call

func myFunction(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

func addOne() func() int { 
	var x int = 0
	return func() int { // defining an anonymous function returns an int value
		x++ // x is within the scope, so it's usable here
		return x
	}
}

func main () {
	fullName := myfunction("Victor", "Yun")
	fmt.Println(fullName)
	secondFullName, err := myFunction("Victor", "Yun") // assign the results to multiple variables
	if err == nil {
		fmt.Println("Caught the error!")
	}
	fmt.Println(secondFullName)

	myAnonymousFunction := addOne()
	fmt.Println(myAnonymousFunction())

}



