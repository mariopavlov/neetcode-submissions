class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        sortedNums = sorted(nums)
        result = []
        base = 0

        while base < len(sortedNums) - 2:

            left = base + 1
            right = len(sortedNums) - 1
            
            while left < right:
                baseElement = sortedNums[base]
                leftElement = sortedNums[left]
                rightElement = sortedNums[right]

                current_sum = baseElement + leftElement + rightElement
                print("Sum: ", current_sum)
                if current_sum == 0:
                    result.append([baseElement, leftElement, rightElement])
                    
                    while left < right and sortedNums[left] == leftElement:
                        left += 1  
                    while left < right and sortedNums[right] == rightElement:
                        right -= 1
                        
                elif current_sum < 0:
                    left += 1
                else:
                    right -= 1
            
            while base < len(sortedNums) - 2 and sortedNums[base] == sortedNums[base + 1]:
                base += 1
            base += 1
        
        return result
        