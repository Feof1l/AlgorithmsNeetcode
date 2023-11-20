package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	/*for i := range s {    //альтернативное удаление из строки символов не относящихся к буквам и цифрам
							//но данный способ значительно медленнее
		if !strings.Contains(alph, string(s[i])) {
			newS = strings.ReplaceAll(newS, string(s[i]), "")

		}
	}*/

	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !(unicode.IsLetter(rune(s[left])) || unicode.IsDigit(rune(s[left]))) {
			left += 1
		}
		for left < right && !(unicode.IsLetter(rune(s[right])) || unicode.IsDigit(rune(s[right]))) {
			right -= 1
		}

		if s[left] != s[right] {
			return false
		} else {
			left += 1
			right -= 1
		}
	}

	return true
}
