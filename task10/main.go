package main

import "fmt"

func sum(nums ...int) int {
	suma := 0
	for i := range nums {
		suma += nums[i]
	}
	return suma
}
func main() {
	a := -1
	summ := 0
	for a != 0 {
		fmt.Scan(&a)
		summ = sum(summ, a)
	}
	fmt.Print(summ)
}
