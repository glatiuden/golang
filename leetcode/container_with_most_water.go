package max_area

func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    most := 0
    for left < right {
        area := min(height[left], height[right]) * (right - left)
        most = max(most, area)
        if (height[left] < height[right]) {
            left++
        } else {
            right--
        }
    }
    return most
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
