package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scanf("%f %f", &b, &a)

	fmt.Printf("%.0f\n", math.Ceil((2.0*a)/b))
}
