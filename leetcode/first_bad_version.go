package first_bad_version

func firstBadVersion(n int) int {
	l := 0
	r := n

	for l <= r {
		mid := l + (r-l)/2
		if l == r {
			return mid
		}

		result := isBadVersion(mid)
		if result {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
