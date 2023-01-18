package insert_intervals

func insert(intervals [][]int, newInterval []int) [][]int {
    var ans [][]int
    pre := newInterval
    for _, curr := range intervals {
        if curr[1] < pre[0] {
            ans = append(ans, curr)
        } else if pre[1] < curr[0] {
            ans = append(ans, pre)
            pre = curr
        } else {
            if curr[0] < pre[0] {
                pre[0] = curr[0]
            }

            if curr[1] > pre[1] {
                pre[1] = curr[1]
            }
        }
    }
    ans = append(ans, pre)
    return ans
}
