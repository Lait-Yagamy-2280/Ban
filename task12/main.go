package main

import "fmt"

func main() {
	N := 0
	fmt.Scan(&N)
	sum := 0
	now := 1
	prev := 0
	for i := 0; i <= N; i++ {
		fmt.Print(sum, " ")
		prev = now
		now = sum
		sum = now + prev
	}
}
