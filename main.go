package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	defer func() {
		if recover := recover(); recover != nil {
			fmt.Println("Recover ", recover)
		}
	}()

	fmt.Println("__ Калькулятор индекса массы тела __")

	for {
		userHeight, userWeight := getUserParameters()

		bodyWeightIndex, err := calculateBodyWeight(userWeight, userHeight)

		if err != nil {
			// fmt.Println(err)
			// continue
			panic("Рост и вес должны быть положительными числами")
		}

		outputResult(bodyWeightIndex)

		isRepeatCalculation := repeatCalculation()

		if !isRepeatCalculation {
			break
		}
	}
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

func calculateBodyWeight(userWeight, userHeight float64) (float64, error) {
	const indexScore float64 = 2

	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("рост и вес должны быть положительными числами")
	}

	bodyWeightIndex := userWeight / math.Pow(userHeight/100, indexScore)

	return bodyWeightIndex, nil
}

func outputResult(bodyWeightIndex float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", bodyWeightIndex)

	switch {
	case bodyWeightIndex < 18.5:
		result += " - недостаточный вес"
	case bodyWeightIndex >= 18.5 && bodyWeightIndex <= 24.9:
		result += " - нормальный вес"
	case bodyWeightIndex >= 25 && bodyWeightIndex <= 29.9:
		result += " - избыточный вес"
	case bodyWeightIndex >= 30:
		result += " - ожирение"
	}

	fmt.Println(result)
}

func repeatCalculation() bool {
	var userChoice string

	fmt.Print("Хотите повторить расчет? (y/n): ")
	fmt.Scanln(&userChoice)

	if userChoice == "y" || userChoice == "Y" {
		return true
	} else {
		return false
	}
}
