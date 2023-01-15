package two_sum

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        j := m[target - curr]
        if (j != 0) {
            return []int{i, j - 1}
        }
        m[curr] = i + 1
    }
    return []int{}
}