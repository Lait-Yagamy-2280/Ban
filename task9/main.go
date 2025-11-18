package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println("Сумма: ", a+b)
	fmt.Println("Разность:", a-b)
	fmt.Println("Произведение:", a*b)
	if b == 0 {
		fmt.Println("Деление на ноль невозможно.")
	} else {
		fmt.Printf("Частное:%.4f\n", float64(a)/float64(b))
	}
}
