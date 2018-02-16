package romanNumerals

import (
	"fmt"
	"sort"
)

var letterMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var romanNumber string
var intNumber int

func IntToRoman(input int) (string, error) {
	if input >= 4000 || input < 1 {
		return "", fmt.Errorf("%d is an invalid value.  Input must be an integer between 1 - 4000", input)
	}

	intNumber = input
	romanNumber = ""
	convertIntToRoman()
	return romanNumber, nil
}

func convertIntToRoman() {
	// Convert map into a slice so it can be sorted
	var values []int
	for value := range letterMap {
		values = append(values, value)
	}

	// Sort it from largest to smallest
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	// Now we can iterate over the slice
	for _, value := range values {
		mapIntToLetters(value)
	}
}

func mapIntToLetters(value int) {
	workingNum := intNumber
	for workingNum > 0 && workingNum >= value && intNumber-value >= 0 {
		workingNum -= value
		romanNumber += letterMap[value]
	}
	intNumber = workingNum
}
