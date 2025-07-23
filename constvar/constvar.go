package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.5

	area := PI * math.Pow(raio, 2)
	fmt.Println(area)
}
