package main

import "fmt"

func main() {
	var name string = "mute"
	var nameAdress *string = &name
	fmt.Println("name(Value) :", name)
	fmt.Println("name(Adress) :", &name)
	fmt.Println("nameAdress (Value) :", *nameAdress)
	fmt.Println("nameAdress (Adress) : ", nameAdress)

	name = "aykyuu"
	fmt.Println("name(Value) :", name)
	fmt.Println("name(Adress) :", &name)
	fmt.Println("nameAdress (Value) :", *nameAdress)
	fmt.Println("nameAdress (Adress) : ", nameAdress)
}
