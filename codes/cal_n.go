package main

import (
	"fmt"
)

func cal(k int) int {
	var i int = 1
	var sum float64 = 0.0
	for true {
		sum += float64(1) / float64(i)
		if sum > float64(k) {
			break
		}
		i++
	}
	return i
}

func main() {
	var k int
	var n int
	fmt.Scanf("%d", &k)
	n = cal(k)
	fmt.Printf("%d", n)
}
