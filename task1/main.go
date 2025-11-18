package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	ans := ""
	if n == 0 {
		ans = "zero"
	} else {
		if n > 0 {
			ans = "positive "
		} else {
			ans = "negative "
		}
		if n%2 == 0 {
			ans += "even"
		} else {
			ans += "odd"
		}
	}
	fmt.Println(ans)
}
