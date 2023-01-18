package product_except_self

func productExceptSelf(nums []int) []int {
    length := len(nums)
    prefix := make([]int, length)
    suffix := make([]int, length)

    prefix_sum := 1
    for i := 0; i < length; i++ {
        prefix[i] = prefix_sum
        prefix_sum *= nums[i]
    }

    suffix_sum := 1
    for i := length - 1; i >= 0; i-- {
        suffix[i] = suffix_sum
        suffix_sum *= nums[i]
    }

    results := make([]int, length)
    for i := 0; i < length; i++ {
        results[i] = prefix[i] * suffix[i] 
    }
	
    return results
}