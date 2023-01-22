package valid_anagram

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
    s_arr := sortString(s)
    t_arr := sortString(t)
    return s_arr == t_arr
}

func sortString(str string) string {
    s := strings.Split(str, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
