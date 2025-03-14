package main

import (
	"fmt"
	"strings"
)

var romanValues = map[rune]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}

func calculateValues(values string) []int {
	var individualValues []int
	for _, digit := range values {
		individualValues = append(individualValues, romanValues[digit])
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

func convertNumbers(romanInput string) int {
	allValues := calculateValues(romanInput)
	signedAllValues := setSumAbst(allValues)
	totalSum := 0
	fmt.Println("All values:", allValues)
	for _, value := range signedAllValues {
		totalSum += value
	}
	return totalSum
}

func isValidRomanNumber(text string) bool {
	if isSyntacticallyCorrect(text) && isGrammaticallyCorrect(text) {
		fmt.Println("Is the number valid? true")
		return true
	}
	fmt.Println("Is the number valid? false")
	return false
}

func isGrammaticallyCorrect(text string) bool {
	threeDigitCounter := 0
	for i, x := range text {
		if i > 0 && rune(text[i-1]) == x {
			threeDigitCounter++
		} else {
			threeDigitCounter = 0
		}
		if i > 2 && rune(text[i-2]) < x && rune(text[i-1]) < x {
			fmt.Println("No more than 3 consecutive identical digits as a subst")
			return false
		}
		if threeDigitCounter == 3 {
			fmt.Println("No more than 3 consecutive identical digits are valid")
			return false
		}
	}
	return true
}

func isSyntacticallyCorrect(text string) bool {
	for _, x := range text {
		if _, exists := romanValues[x]; !exists {
			fmt.Println("Invalid letter found:", string(x))
			return false
		}
	}
	return true
}

func main() {
	var isValidNumber bool
	var upperInput string

	for !isValidNumber {
		var inputText string
		fmt.Print("Give a valid roman notation number: ")
		fmt.Scanln(&inputText)
		upperInput = strings.ToUpper(inputText)
		fmt.Println("Input:", upperInput)
		isValidNumber = isValidRomanNumber(upperInput)
	}

	fmt.Println(convertNumbers(upperInput))
}
