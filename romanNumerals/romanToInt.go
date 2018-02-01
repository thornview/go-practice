// RomanToInt converts a roman numeral, provided as a command argument,
// into its integer equivalent.  It assumes the value will be less than 4,000
package romanNumerals

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

func romanToInt() {
	romanNum := strings.ToUpper(os.Args[1])
	intValues := convertRomanToInt(romanNum)
	_, err := isValidRomanNumeral(intValues)
	if err != nil {
		fmt.Printf("Error in %s: %v", romanNum, err)
	} else {
		sumTotal := calculateTotal(intValues)
		fmt.Printf("Total: %d\n", sumTotal)
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

func isValidRomanNumeral(numberSequence []int) (bool, error) { // Using the bool return in the tests
	for i := 0; i < len(numberSequence); i++ {
		if numberSequence[i] == 0 {
			return false, fmt.Errorf("contains invalid character")
		}
		if i > 0 {
			thisNum := numberSequence[i]
			prevNum := numberSequence[i-1]
			if thisNum > prevNum && !(prevNum*10 == thisNum || prevNum*5 == thisNum) {
				// Cannot have higher number after lower number except multiples of 10 and 5
				return false, fmt.Errorf("letters out of order")
			}
			if i > 1 {
				penultimateNum := numberSequence[i-2]
				if thisNum > prevNum && prevNum == penultimateNum {
					// Cannot have BBA series (ex. iiv, xxl)
					return false, fmt.Errorf("contains high number after a chain of low numbers")
				}
			}
		}

		// TODO Other Tests:
		// Some BAB sequences are invalide (viv, ixi, etc.) but some aren't (mcm, cxc).  Not sure of rule
		// cannot have more than 4 of any number
	}
	return true, nil
}
