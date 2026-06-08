class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        longest_consecutive = 0

        nums_set = set(nums)
        sequence_starters = []

        for num in nums:
            if not (num - 1) in nums_set:
                sequence_starters.append(num)
        
        for num in sequence_starters:
            temp_consecutive = 1
            lookup = num

            while (lookup + 1) in nums_set:
                temp_consecutive += 1
                lookup += 1
            
            if temp_consecutive > longest_consecutive:
                longest_consecutive = temp_consecutive
        
        return longest_consecutive