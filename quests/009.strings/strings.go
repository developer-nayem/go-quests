package strings

import (
	"unicode/utf8"
)

type TextStats struct {
	ByteLength int
	RuneCount  int
}

func AnalyzeText(s string) TextStats {
	// TODO: implement
	// Read README.md for the instructions
	byteLength := len(s)
	runeCount := utf8.RuneCountInString(s)
	return TextStats{
		byteLength,
		runeCount,
	}
}

func RuneFrequencies(s string) map[rune]int {
	// TODO: implement
	// Read README.md for the instructions
	myMap := make(map[rune]int)

	for _,value := range s {
		_,exists := myMap[value]
		if !exists {
			myMap[value] = 1
		}else {
			myMap[value] += 1
		}
	}
	return myMap
}

func FirstRunePosition(s string, target rune) int {
	// TODO: implement
	// Read README.md for the instructions

	for i, value := range s {
		if value == target {
			return i
		}
	}
	return -1
}

func ExtractRunes(s string) []rune {
	// TODO: implement
	// Read README.md for the instructions

	result := []rune{}

	for _, value := range s {
		result = append(result, value)
	}

	return result 
	
}

func HasOnlyASCII(s string) bool {
	// TODO: implement
	// Read README.md for the instructions
	if len(s) == 0 {
		return true 
	}

	for _, value := range s {
		if int(value) > 127 {
			return false 
		}
	}


	return true
}
