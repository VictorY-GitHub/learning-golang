package main

import (
	"fmt"
)

type Employee struct {
	Name string
}

func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

func main () {
	var emp Employee
	emp.Name = "Victor Yun" // difference between function and methods, methods act on an object e.g. e.SOMETHING
	emp.UpdateName("funmblr")
	emp.PrintName()
}