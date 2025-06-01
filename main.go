package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight, userWeight := getUserParameters()

	if userHeight <= 0 || userWeight <= 0 {
		fmt.Println("Ошибка: рост и вес должны быть положительными числами")

		return
	}

	bodyWeightIndex := calculateBodyWeight(userWeight, userHeight)

	outputResult(bodyWeightIndex)
}

func getUserParameters() (float64, float64) {
	var userHeight float64
	var userWeight float64

	fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scanln(&userHeight)
	fmt.Print("Введите ваш вес в килограммах: ")
	fmt.Scanln(&userWeight)

	return userHeight, userWeight
}

func calculateBodyWeight(userWeight, userHeight float64) float64 {
	const indexScore float64 = 2
	bodyWeightIndex := userWeight / math.Pow(userHeight/100, indexScore)

	return bodyWeightIndex
}

func outputResult(bodyWeightIndex float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", bodyWeightIndex)
	fmt.Print(result)
}
