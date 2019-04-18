package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s\n", &str)

	if isPalindromic(str) {
		fmt.Printf("%s is palindromic", str)
	}

	if isPalindromicByChinese(str) {
		fmt.Printf("%s is palindromicByChinese", str)
	}
}

func isPalindromic(s string) bool {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		switch {
		case i == sLen:
			break
		case s[i] != s[sLen-1-i]:
			return false
		}
	}
	return true
}

func isPalindromicByChinese(s string) bool {
	// convert to rune
	runes := []rune(s)
	rLen := len(runes)

	for i := 0; i < rLen/2; i++ {
		switch {
		case i == rLen:
			break
		case runes[i] != runes[rLen-1-i]:
			return false
		}
	}
	return true
}