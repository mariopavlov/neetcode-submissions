func removeDuplicates(nums []int) int {
    n := len(nums)
    l := 0
    r := 0

    for r < n {
        nums[l] = nums[r]
        
        for r < n && nums[r] == nums[l] {
            r++
        }

        l++
    }

    return l
}
