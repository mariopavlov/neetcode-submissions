func removeDuplicates(nums []int) int {
    left := 1

    for right := 1; right < len(nums);right++ {

        if nums[right] == nums[right - 1] {
            continue
        } 

        nums[left] = nums[right]
        left++
    }

    return left
}
