package main

import "fmt"

type employee struct {
	firstName, lastName string
}

func (e employee) fullName() string {
	return e.firstName + " " + e.lastName
}

func main() {
	e := employee{
		firstName: "muhamad",
		lastName:  "ismail",
	}
	fmt.Println(e.fullName())
}
