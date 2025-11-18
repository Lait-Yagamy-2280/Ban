package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	F := c*9/5 + 32
	return float64(F)
}

func (c *Celsius) Add(d Celsius) {
	*c += d

}
func main() {
	var a, b Celsius
	fmt.Scan(&a, &b)
	fmt.Printf("First temp: %.2f\n", a.ToFahrenheit())
	fmt.Printf("Second temp: %.2f\n", b.ToFahrenheit())
	a.Add(b)
	fmt.Printf("Sum temp's: %.2f\n", float64(a))
}
