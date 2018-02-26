package wordGame

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetLetterSet() {
	numLetters := rand.Intn(4) + 4
	regLetters := getRandomLetter(numLetters - 1)
	vowels := getVowels(1)
	fmt.Printf("Selected %v letters: \n Regular set: %v \n Vowel: %v\n", numLetters, regLetters, vowels)
}

func getRandomLetter(numLetters int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	return getRandomLettersFromSet(alphabet, numLetters)
}

func getVowels(numLetters int) string {
	vowels := "aeiou"
	return getRandomLettersFromSet(vowels, numLetters)
}

func getRandomLettersFromSet(letterSet string, numLetters int) string {
	selectLetters := make([]byte, numLetters)
	for i := range selectLetters {
		selectLetters[i] = letterSet[rand.Intn(len(letterSet))]
	}
	return string(selectLetters)
}
