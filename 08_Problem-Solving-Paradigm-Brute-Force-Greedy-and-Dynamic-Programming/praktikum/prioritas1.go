package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binaryRepresentation(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	return ans
}

func main() {
	n1 := 2
	n2 := 3

	for _, n := range []int{n1, n2} {
		result := binaryRepresentation(n)
		fmt.Printf("Input: n = %d\n", n)
		fmt.Printf("Output: [%s]\n", strings.Join(result, ", "))
	}
}
