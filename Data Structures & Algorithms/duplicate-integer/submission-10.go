func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{}, len(nums))

	for _, elem := range nums {
		if _, ok := seen[elem]; ok {
			return true
		}

		seen[elem] = struct{}{}
	}

	return false
}
