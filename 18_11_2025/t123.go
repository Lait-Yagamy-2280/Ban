package main

import (
	"fmt"
	"time"
)

func main() {
	//task 1
	days := [...]string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}
	n := int(time.Now().Weekday())-1
	slice1 := days[n:]
	fmt.Printf("Massive: %v %v\n", len(days), cap(days))
	fmt.Printf("Massive: %v %v\n", len(slice1), cap(slice1))

	//task2
	for i := range slice1 {
		fmt.Print(days[i], " ")
	}
	fmt.Print("\n", days[n], "\n")
	for i := range slice1 {
		fmt.Print(slice1[i], " ")
	}
	fmt.Println()

	//task3
	slice2 := make([]string, len(slice1), len(slice1)+7)
	copy(slice2, slice1)
	slice2 = append(slice2, "Пятница", "Пятница", "Пятница", "Пятница", "Пятница", "Пятница", "Пятница")
	for i := range slice2 {
		fmt.Print(slice2[i], " ")
	}
	fmt.Println()
}

