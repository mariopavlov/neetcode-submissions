class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
         
         # Solution will use SET
         # Pros: constant time to add a number, constant time to check if number part of set
         # Cons: memory overhead of using a set
         # Time-Complexity: O(n)
         # Time-Complexity Set Operation (AVG): O(1)
         # Time-Complexity Set Operation (Worst-Case): O(n), rare
         # Space-Complexity: O(n)

         # 1. Initialize a SET, set()
         unique_numbers = set()
         
         # 2. Iterate over the input list, add each number in the set
         for num in nums:
            # 2.1 If num is in set
            if num in unique_numbers:
                return True
            
            # 2.2 add to the set
            unique_numbers.add(num)
         
         # 3. All numbers are unique
         return False
