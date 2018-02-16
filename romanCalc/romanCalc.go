package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/thornview/romanNumerals"
)

func main() {
	showMainMenu()
}

func showMainMenu() {
	fmt.Println("[R]oman to int, [I]nt to Roman, [C]alculator, [Q]uit? ")
	result, err := getRune()
	if err != nil {
		showError()
	}
	choice := unicode.ToUpper(result)
	switch choice {
	case 'R':
		runRomanToInt()
		break
	case 'I':
		runIntToRoman()
		break
	case 'C':
		runCalc()
	case 'Q':
		sayGoodbye()
		return
	default:
		showError()
	}
	showMainMenu()
}

func showError() {
	fmt.Printf("\nThat was not a valid input. \nTry Again.\n\n")
	showMainMenu()
}

func getRune() (rune, error) {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadRune()
	return result, err
}

func getLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getNumber() (int, error) {
	input := getLine()
	number, err := strconv.Atoi(input)
	return number, err
}

func runRomanToInt() {
	fmt.Printf("Enter Roman Numeral: ")
	romanNum := getLine()
	romanNumerals.RomanToInt(romanNum)
}

func runIntToRoman() {
	fmt.Printf("Enter Integer: ")
	number, err := getNumber()
	if err != nil {
		showError()
		return
	}
	romanNum, romanErr := romanNumerals.IntToRoman(number)
	if romanErr != nil {
		fmt.Printf("%v", romanErr)
	}
	fmt.Printf("\n\n%v = %s\n\n", number, romanNum)

}

func runCalc() {
	fmt.Printf("Coming soon ... or not")
}

func sayGoodbye() {
	fmt.Printf("\n\nGoodbye\n\n")
}
