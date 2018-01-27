package main

import (
	"fmt"
	"os"
	"strings"
)

var numberMap = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func main() {
	romanNum := strings.ToUpper(os.Args[1])
	intValues := convertRomanToInt(romanNum)
	if isValidRomanNumeral(intValues) {
		sumTotal := calculateTotal(intValues)
		fmt.Printf("Total: %d\n", sumTotal)
	} else {
		errorMessage := "Hmm ... that doesn't look like a Roman Numeral to me."
		fmt.Printf("%s\n", errorMessage)
	}
}

func convertRomanToInt(romanNum string) []int {
	var values []int
	for _, rune := range romanNum { // The _ is a place holder for the index value, which we do not use so cannot assign to a variable
		value := numberMap[string(rune)] // Note: if the rune provided is not a key in the map, the default int value of 0 is returned
		values = append(values, value)
	}
	return values
}

func calculateTotal(values []int) int {
	var sumTotal int
	for i := 0; i < len(values); i++ {
		if i > 0 && values[i] > values[i-1] {
			sumTotal += (-2 * values[i-1])
		}
		sumTotal += values[i]
	}
	return sumTotal
}

func isValidRomanNumeral(numberSequence []int) bool {
	for i := 0; i < len(numberSequence); i++ {
		if i > 0 {
			thisNum := numberSequence[i]
			prevNum := numberSequence[i-1]
			if thisNum == 0 {
				// Letter is not valid for Roman numerals
				return false
			} else if thisNum > prevNum && !(prevNum*10 == thisNum || prevNum*5 == thisNum) {
				// Cannot have higher number after lower number except multiples of 10 and 5
				return false
			}
			if i > 2 {
				// Cannot have ABA series (ex. ivi, xcx)
				penultimateNum := numberSequence[i-2]
				if thisNum == penultimateNum && thisNum != prevNum {
					return false
				}
			}
		} else if i == 0 && numberSequence[i] == 0 {
			return false
		}
	}
	return true
}
