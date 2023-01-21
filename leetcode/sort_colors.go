package sort_colors

func sortColors(nums []int)  {
    start := 0
    middle := 0
    end := len(nums) - 1
    for middle <= end {
        curr := nums[middle]
        if curr == 0 {
            swap(nums, start, middle)
            start++
            middle++
        } else if curr == 2 {
            swap(nums, middle, end)
            end--
        } else {
            middle++
        }
    }
}

func swap(nums []int, i, j int) {
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}
