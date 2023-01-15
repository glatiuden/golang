package contains_duplicate

func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, num := range nums {
        result := m[num]
        if (result == true) {
            return true
        }
        m[num] = true
    }
    return false
}
