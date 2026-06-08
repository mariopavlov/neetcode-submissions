class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        index = 1
        for a in nums:
            
            for b in nums[index:]:
                print("a = ", a)
                print("b = ", b)
                if a == b:
                    return True
            index += 1
        
        return False

         