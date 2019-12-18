package main

import ("fmt")

func main() {
	var daysOfWeek []string
	daysOfWeek = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekdays := daysOfWeek[0:5]
	fmt.Println(weekdays) // sliced from index 0 to 5

	instagramFollowers := map[string]int { // representation of hash tables
		"funmblr": 119000,
		"thefandomshelter": 400000,
	}
	fmt.Println(instagramFollowers["funmblr"]) // 119000

	type Person struct {
		name string
		age int
	}
	var victorYun Person = Person{name: "Victor", age: 19}
	victorYun.age = 18
	fmt.Println(victorYun.name) // Victor
	fmt.Println(victorYun.age) // 18
}
