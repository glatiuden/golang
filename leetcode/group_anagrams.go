package group_anagarams

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    for _, str := range strs {
        sorted := sortString(str)
        m[sorted] = append(m[sorted], str)
    }
    v := make([][]string, 0, len(m))
    for  _, value := range m {
        v = append(v, value)
    }
    return v
}

func sortString(str string) string {
    s := strings.Split(str, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
