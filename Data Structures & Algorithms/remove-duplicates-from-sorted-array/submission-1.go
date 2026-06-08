func removeDuplicates(nums []int) int {
    cur := 0
    counter := 1

    for i := 1; i < len(nums); i++ {
        fmt.Println(nums[i])

        if nums[i] == nums[cur] {
            continue
        } 

         cur++
         nums[cur] = nums[i]
        counter++
    }

    return counter
}
