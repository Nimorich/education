package main

import (
	"fmt"
	"math"
	"errors"
)

const IMTPower = 2

func main() {
		fmt.Println("__ Калькулятор индекса массы тела __ ")
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
		// 	fmt.Println(err)
		// 	continue
		panic(err)
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}
func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("Не указан вес или высота")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}
func outputResult(imt float64) {
	result := fmt.Sprintf("Ваши индекс массы тела: %.0f", imt)
	fmt.Print(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("у вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}
func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост; ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес; ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчёт (y/n):")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
