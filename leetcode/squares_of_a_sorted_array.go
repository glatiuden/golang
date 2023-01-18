package square_of_a_sorted_array

import "sort"

func sortedSquares(nums []int) []int {
    squared := []int{}
    for _, num := range nums {
        squared = append(squared, num*num)
    }
    sort.Ints(squared)
    return squared
}