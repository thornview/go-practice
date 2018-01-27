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
	sumTotal := calculateTotal(intValues)
	fmt.Printf("Total: %d", sumTotal)
}

func convertRomanToInt(romanNum string) []int {
	var values []int
	for _, rune := range romanNum {
		value := numberMap[string(rune)]
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
