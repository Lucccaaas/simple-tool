package common

import "unicode"

func UpperCaseFirst(key string) string {
	runes := []rune(key)
	if (runes[0] >= 'a') && (runes[0] <= 'z') {
		runes[0] = unicode.ToUpper(runes[0])
	}
	result := string(runes)
	return result
}

func LowerCaseFirst(key string) string {
	runes := []rune(key)
	if (runes[0] >= 'A') && (runes[0] <= 'Z') {
		runes[0] = unicode.ToLower(runes[0])
	}
	result := string(runes)
	return result
}
