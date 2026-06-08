func hasDuplicate(nums []int) bool {
    seen := map[int]struct{}{}

	for _, elem := range nums {
		_, ok := seen[elem]
		if ok {
			return true
		}

		seen[elem] = struct{}{}
	}

	return false
}
