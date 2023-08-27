package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// cara pertama

	var person1 student
	person1.firstName = "rayan"
	person1.lastName = "ahmad"
	person1.age = 20
	fmt.Println(person1)

	//cara kedua
	var person2 = student{
		firstName: "jamuadil",
		lastName:  "Abdullah",
		age:       22,
	}
	fmt.Println(person2)

	// Cara ketiga
	person3 := student{"malika", "danti", 50}
	fmt.Println(person3)
}
