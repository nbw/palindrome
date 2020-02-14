package palindrome

import (
	"testing"
)

func TestIsPalindromeTrue(t *testing.T) {
	tests := []string{
		"racecar",
		"race car",
		"race .!@#$%!@#car!@#!@#!@",
		"raCecar",
		"raC !#*ecar",
	}

	var result string

	for i := 0; i < len(tests); i++ {
		result = IsPalindrome(tests[i])
		if result != "A palindrome" {
			t.Error(
				"For", tests[i],
				"expected", "A palindrome",
			)
		}
	}
}

func TestIsPalindromeFalse(t *testing.T) {
	tests := []string{
		"abcdefg",
		"ab cdefg",
		"ab cde!fg !@#(*&!@",
	}

	var result string

	for i := 0; i < len(tests); i++ {
		result = IsPalindrome(tests[i])
		if result != "Not a palindrome" {
			t.Error(
				"For", tests[i],
				"expected", "Not a palindrome",
			)
		}
	}
}
