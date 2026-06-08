class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        items = {}

        for idx, num in enumerate(nums):
            diff = target - num

            if diff in items:
                return [items[diff], idx]

            items[num] = idx

        return []