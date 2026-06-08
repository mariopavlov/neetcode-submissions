class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # 2nd Atempt to solve the problem
        # with a little bit of help I saw that I have to use dictionary

        items = {}

        for idx, i in enumerate(nums):
            difference = target - i

            if (difference in items):
                return [items[difference], idx]
            
            items[i] = idx


        return [0,0]


        