package find_the_town_judge

func findJudge(n int, trust [][]int) int {
    arr := make([]int, n)
    for _, curr := range trust {
        arr[curr[0] - 1]--
        arr[curr[1] - 1]++
    }

    for i, val := range arr {
        if val == n - 1 {
            return i + 1
        }
    }

    return -1
}