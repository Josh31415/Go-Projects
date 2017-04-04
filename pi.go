package main

import "fmt"

func main() {

	var pi float64

	step := 1000000000
	
	x := 3.0
	y := -1.0
	pi = 4.0

	for i := 0; i < step; i++ {
			
		pi =  pi + (4.0 / x * y)
		x += 2.0
		y = -y
	}
	fmt.Println("pi = ", pi)
}
