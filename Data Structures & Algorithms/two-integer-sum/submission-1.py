class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # 1st solution that comes to my mind
        # two pointers that will read from the array and check for the target

        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]
        
        return [0,0]