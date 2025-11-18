package main

import "fmt"

func main() {
	score := 0
	fmt.Scan(&score)
	if (score <= 100) && (score >= 0) {
		switch score / 10 {
		case 9, 10:
			fmt.Println("A")
		case 8:
			fmt.Println("B")
		case 7:
			fmt.Println("C")
		case 6:
			fmt.Println("D")
		case 0, 1, 2, 3, 4, 5:
			fmt.Println("F")
		}
	} else {
		fmt.Println("invalid")
	}
}
