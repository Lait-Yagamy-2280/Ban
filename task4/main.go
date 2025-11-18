package main

import "fmt"

func swap(a, b *int) {
	c := *b
	*b = *a
	*a = c
}
func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	swap(&a, &b)
	fmt.Println(a, b)
}
