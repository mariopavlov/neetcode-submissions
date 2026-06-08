
func findMaxConsecutiveOnes(nums []int) int {
	cur := 0

	local := 0
	for _, i := range nums {
		if i == 1 {
			local += 1
		} else {
			cur = max(cur, local)
			local = 0
		}
	}

	return max(cur, local)
}
