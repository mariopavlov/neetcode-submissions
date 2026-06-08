class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        lookFor = 0

        result = []

        a = 0
        

        while a < len(nums) - 2:
            b = a + 1
            c = len(nums) - 1
            
            while b < c:
                elementA = nums[a]
                elementB = nums[b]
                elementC = nums[c]

                sum = elementA + elementB + elementC
                
                if sum == lookFor:
                    result.append([elementA, elementB, elementC])

                    while b < c and nums[b] == elementB:
                        b += 1

                    while b < c and nums[c] == elementC:
                        c -= 1

                elif sum < lookFor:
                    b += 1
                    
                elif sum > lookFor:
                    c -= 1

            while a < len(nums) - 2 and nums[a] == elementA:
                a +=1

        return result