package longest_palindrome

import "strings"

func longestPalindrome(s string) int {
    if len(s) == 1 { return 1 }
    _map := make(map[string]int)

    str := strings.Split(s, "")
    for _, char := range str {
        _map[char]++
    }

    odd := 0
    for _, count := range _map {
        if count % 2 != 0 {
            odd++
        }
    }

    if odd == 0 {
        return len(s)
    }

    if len(s) < odd {
        return 1
    }

    return len(s) + 1 - odd
}