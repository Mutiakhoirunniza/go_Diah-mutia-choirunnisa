package main

import "fmt"

func SimpleEquations(a, b, c int) {
	for x := 1; x <= a; x++ { // Fungsi ini mencoba semua kemungkinan nilai x, y, dan z dalam rentang 1 hingga 10000.
		for y := 1; y <= a; y++ {
			for z := 1; z <= a; z++ {
				if x != y && y != z && z != x && x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Printf("%d %d %d\n", x, y, z)
					return
				}
			}
		}
	}
	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
