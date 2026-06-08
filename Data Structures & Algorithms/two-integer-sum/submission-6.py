class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        processed = {}
        
        for idx, num in enumerate(nums):
            diff = target - num
            if diff in processed:
                return [processed.get(diff), idx]
        
            processed[num] = idx

        return []