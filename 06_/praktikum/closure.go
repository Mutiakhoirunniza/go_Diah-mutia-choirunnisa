package main

import "fmt"

func newCounter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}

}

func main() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
