package valid_palindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
    m := regexp.MustCompile("[^a-zA-Z0-9]")
    str := strings.ToLower(m.ReplaceAllString(s, ""))
    reversed := strings.ToLower(reverse(str))
    return reversed == str
}

func reverse(str string) string {
    var result string
    for _, v := range str {
        result = string(v) + result
    }
    return result
}