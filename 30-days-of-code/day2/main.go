package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		mealCost   float64
		tipPercent int
		taxPercent int
	)
	fmt.Scanf("%f\n%d\n%d\n", &mealCost, &tipPercent, &taxPercent)
	tip := mealCost * float64(tipPercent) / 100
	tax := mealCost * float64(taxPercent) / 100
	totalCost := mealCost + tip + tax
	x, y := math.Modf(totalCost)
	if y >= 0.5 {
		totalCost = x + 1
		fmt.Printf("The total meal cost is %v dollars.\n", totalCost)
	} else {
		totalCost = x
		fmt.Printf("The total meal cost is %v dollars.\n", totalCost)
	}

	/*
		// Изящный вариант
		fmt.Scanln(&mealCost)
		fmt.Scanln(&tipPercent)
		fmt.Scanln(&taxPercent)
		tip := mealCost * (float64(tipPercent) / 100.0)
		tax := mealCost * (float64(taxPercent) / 100.0)
		totalCost := mealCost + tip + tax
		fmt.Printf("The total meal cost is %.0f dollars.", totalCost)
	*/
}
