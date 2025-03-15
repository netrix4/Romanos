package main

import (
	"fmt"
	"strings"
	"strconv"
)

var romanValues = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50,
	"C": 100, "D": 500, "M": 1000,
}

func calculateValues(values string) []int {
	var individualValues []int
	for _, digit := range values {
		individualValues = append(individualValues, romanValues[string(digit)])
	}
	return individualValues
}

func setSumAbst(arabicValues []int) []int {
	for i := 1; i < len(arabicValues); i++ {
		if arabicValues[i-1] < arabicValues[i] {
			arabicValues[i-1] *= -1
		}
	}
	return arabicValues
}

func convertUnits(digit int) string {
	switch digit {
	case 9:
		return "IX"
	case 4:
		return "IV"
	case 5, 6, 7, 8:
		return "V" + strings.Repeat("I", digit-5)
	default:
		return strings.Repeat("I", digit)
	}
}

func convertTens(digit int) string {
	switch digit {
	case 9:
		return "XC"
	case 4:
		return "XL"
	case 5, 6, 7, 8:
		return "L" + strings.Repeat("X", digit-5)
	default:
		return strings.Repeat("X", digit)
	}
}

func convertHundreds(digit int) string {
	switch digit {
	case 9:
		return "CM"
	case 4:
		return "CD"
	case 5, 6, 7, 8:
		return "D" + strings.Repeat("C", digit-5)
	default:
		return strings.Repeat("C", digit)
	}
}

func convertThousands(digit int) string {
	return strings.Repeat("M", digit)
}

func convertArabicToRoman(number int) string {
	thousands := number / 1000
	hundreds := (number % 1000) / 100
	tens := (number % 100) / 10
	units := number % 10

	return convertThousands(thousands) +
		convertHundreds(hundreds) +
		convertTens(tens) +
		convertUnits(units)
}

func convertNumbers(romanInput string) int {
	allValues := calculateValues(romanInput)
	signedAllValues := setSumAbst(allValues)
	totalSum := 0
	for _, value := range signedAllValues {
		totalSum += value
	}
	return totalSum
}

func isValidRomanNumber(text string) bool {
	return isSyntacticallyCorrect(text) && isGrammaticallyCorrect(text)
}

func isValidArabicNumber(numberStr string) bool {
	_, err := strconv.Atoi(numberStr)
	return err == nil
}

func isGrammaticallyCorrect(text string) bool {
	threeDigitCounter := 0
	for i := 0; i < len(text); i++ {
		if i > 0 && text[i-1] == text[i] {
			threeDigitCounter++
		} else {
			threeDigitCounter = 0
		}
		if threeDigitCounter == 3 {
			return false
		}
	}
	return true
}

func isSyntacticallyCorrect(text string) bool {
	for _, char := range text {
		if _, exists := romanValues[string(char)]; !exists {
			return false
		}
	}
	return true
}

func romanToArabic() {
	var inputText string
	for {
		fmt.Print("Give a valid roman notation number: ")
		fmt.Scanln(&inputText)
		inputText = strings.ToUpper(inputText)
		if isValidRomanNumber(inputText) {
			break
		}
	}
	fmt.Printf("Your number is: %d\n", convertNumbers(inputText))
}

func arabicToRoman() {
	var inputText string
	for {
		fmt.Print("Give a valid arabic notation number: ")
		fmt.Scanln(&inputText)
		if isValidArabicNumber(inputText) {
			break
		}
	}
	number, _ := strconv.Atoi(inputText)
	fmt.Printf("Your number is: %s\n", convertArabicToRoman(number))
}

func main() {
	var conversionOption string
	fmt.Println("(1) Roman to Arabic: MCMXCVII -> 1997")
	fmt.Println("(2) Arabic to Roman: 1997 -> MCMXCVII")
	fmt.Print("Select an option: ")
	fmt.Scanln(&conversionOption)

	if conversionOption == "1" {
		romanToArabic()
	} else if conversionOption == "2" {
		arabicToRoman()
	}
}
