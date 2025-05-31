package main

import (
	"fmt"
	"math"
)

func main() {
	const indexScore float64 = 2
	userHeight := 1.8
	userWeight := 100.0
	BodyWeightIndex := userWeight / math.Pow(userHeight, indexScore)

	fmt.Print(BodyWeightIndex)
}
