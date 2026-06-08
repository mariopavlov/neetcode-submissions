class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        processed = set()

        for num in nums:
            if num not in processed:
                processed.add(num)
            else:
                return True
        
        return False
         