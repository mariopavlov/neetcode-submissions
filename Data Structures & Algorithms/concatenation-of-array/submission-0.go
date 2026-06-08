func getConcatenation(nums []int) []int {
	output := make([]int, 0, len(nums)*2)
	output = append(output, nums...)
	output = append(output, nums...)

	return output
}
