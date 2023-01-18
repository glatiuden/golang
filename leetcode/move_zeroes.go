package move_zeroes

func moveZeroes(nums []int)  {
    swap := 0
    for i, num := range nums {
        if num != 0 {
            temp := nums[swap]
            nums[swap] = num
            nums[i] = temp
            swap += 1
        }
    }
}