package main

import "fmt"

const PI = 3.14159

func main() {
	r := 0.0
	fmt.Scan(&r)
	s := PI * r * r
	fmt.Println("Area:", s)
}
