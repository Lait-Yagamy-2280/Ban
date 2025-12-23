package main

import (
	"fmt"
	"math"
)

func safe_del(a, b int) (string, bool) {
	flag := true
	ans := ""
	if b == 0 {
		flag = false
	}
	if flag {
		c := math.Pow(10, float64(2))
		ansif := (float64(math.Round(float64(float64(a)/float64(b))*c) / c))
		ans = fmt.Sprint(ansif)
	} else {
		ans = "Error"
	}
	return ans, flag
}
func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	an, f := safe_del(a, b)
	if f {
		fmt.Println("Деление прошло штатно, ответ:", an)
	} else {
		fmt.Println("Возникла ошибка при делении.")
	}
}
