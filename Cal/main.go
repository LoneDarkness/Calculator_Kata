package main

import (
	"fmt"
	"strconv"
)

func isArabic(varString1, varString2 string) bool {
	trueArabicMap := map[string]bool{
		"1":  true,
		"2":  true,
		"3":  true,
		"4":  true,
		"5":  true,
		"6":  true,
		"7":  true,
		"8":  true,
		"9":  true,
		"10": true,
	}
	return trueArabicMap[varString1] && trueArabicMap[varString2]
}

func isRoman(varString1, varString2 string) bool {
	trueRomanMap := map[string]bool{
		"I":    true,
		"II":   true,
		"III":  true,
		"IV":   true,
		"V":    true,
		"VI":   true,
		"VII":  true,
		"VIII": true,
		"IX":   true,
		"X":    true,
	}
	return trueRomanMap[varString1] && trueRomanMap[varString2]
}

func calArabic(var1, var2 int, sign string) (result int) {
	switch sign {
	case "+":
		result = var1 + var2
	case "-":
		result = var1 - var2
	case "*":
		result = var1 * var2
	case "/":
		result = var1 / var2
	default:
		panic("Error")
	}

	fmt.Println(result)
	return result
}

func calRoman(var1, var2 int, sign string) (resultInt int) {
	switch sign {
	case "+":
		resultInt = var1 + var2
	case "-":
		resultInt = var1 - var2
	case "*":
		resultInt = var1 * var2
	case "/":
		resultInt = var1 / var2
	default:
		panic("Error")
	}
	if resultInt < 0 {
		panic("Ошибка: Значение ниже 0. В римской системе нет отрицательных чисел")
	}
	return resultInt
}

func intToRoman(resultInt int) string {
	symbol := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""

	for resultInt > 0 {
		for i := range value {
			if resultInt >= value[i] {
				result += symbol[i]
				resultInt -= value[i]
				break
			}
		}
	}
	fmt.Println(result)
	return result

}

func arabic(varString1, varString2 string) (var1, var2 int) {
	var1, err1 := strconv.Atoi(varString1)
	if err1 != nil {
		panic("Error")
	}

	var2, err2 := strconv.Atoi(varString2)
	if err2 != nil {
		panic("Error")
	}
	return var1, var2
}

func roman(varString1, varString2 string) (var1, var2 int) {
	romanToArabic := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	var1 = romanToArabic[varString1]
	var2 = romanToArabic[varString2]
	return var1, var2
}

func main() {

	var varString1, varString2, sign string
	n, err := fmt.Scanln(&varString1, &sign, &varString2)
	if err != nil || n != 3 {
		fmt.Println("Ошибка: введите ровно две целочисленные переменные.")
		return
	}

	if isArabic(varString1, varString2) == true {
		var1, var2 := arabic(varString1, varString2)
		calArabic(var1, var2, sign)
	} else if isRoman(varString1, varString2) == true {
		var1, var2 := roman(varString1, varString2)
		resultInt := calRoman(var1, var2, sign)
		intToRoman(resultInt)
	} else {
		panic("Error")
	}
}
