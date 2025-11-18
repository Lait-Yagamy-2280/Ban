package main

import "fmt"

func divmod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("Деление на ноль.")
	} else {
		return a / b, a % b, nil
	}
}
func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	ans1, ans2, e := divmod(a, b)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(ans1, ans2)
	}
}
