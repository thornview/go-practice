
// RomanToInt converts a roman numeral, provided as a command argument, 
// into its integer equivalent.  It assumes the value will be less than 4,000
// TODO: Write tests
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
var errorMessage string

func main() {
	romanNum := strings.ToUpper(os.Args[1])
	intValues := convertRomanToInt(romanNum)
	if isValidRomanNumeral(intValues) {
		sumTotal := calculateTotal(intValues)
		fmt.Printf("Total: %d\n", sumTotal)
	} else {
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
				errorMessage = "Oops, you added a character that doesn't belong in a Roman Numeral."
				return false
			} else if thisNum > prevNum && !(prevNum*10 == thisNum || prevNum*5 == thisNum) {
				// Cannot have higher number after lower number except multiples of 10 and 5
				errorMessage = "It looks like you have some of those letters out of order."
				return false
			}
			if i > 1 {
				penultimateNum := numberSequence[i-2]
				if thisNum == penultimateNum && thisNum != prevNum {
					// Cannot have BAB series (ex. ivi, xcx)
					errorMessage = "Hmm ... you can't have a low-high-low sequence like that."
					return false
				} else if thisNum > prevNum && prevNum == penultimateNum {
					// Cannot have BBA series (ex. iiv, xxl)
					errorMessage = "Hey, I found a high number after two low numbers.  I'm confused."
					return false
				}
			}
		} else if i == 0 && numberSequence[i] == 0 {
			errorMessage = "Wow, that got off to a bad start.  That first character isn't a Roman Numeral."
			return false
		}
		// cannot have more than 4 of any number
		// TODO write a slice test to verify
	}
	return true
}
