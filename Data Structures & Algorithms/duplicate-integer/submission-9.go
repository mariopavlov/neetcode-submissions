func hasDuplicate(nums []int) bool {
    seen := map[int]struct{}{}

	for _, elem := range nums {
		if _, ok := seen[elem]; ok {
			return true
		}

		seen[elem] = struct{}{}
	}

	return false
}
