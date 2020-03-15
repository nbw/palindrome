package palindrome

import (
	"log"
	"regexp"
	"strings"
)

func IsPalindrome(input string) string {
	if isPalindrome(input) == true {
		return "A palindrome"
	} else {
		return "Not a palindrome"
	}
}

func isPalindrome(input string) bool {
	var sanitizedInput string
	sanitizedInput = removePunctuation(input)
	sanitizedInput = strings.ToLower(sanitizedInput)

	length := len(sanitizedInput)
	for i := 0; i < length/2; i++ {
		if sanitizedInput[i] != sanitizedInput[length-1-i] {
			return false
		}
	}

	return true
}

func removePunctuation(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(input, "")
}
