class Solution:
    def isPalindrome(self, s: str) -> bool:
        left = 0
        right = len(s) - 1
        
        while left < right:
            leftChar = s[left].lower()
            rightChar = s[right].lower()
            
            if not leftChar.isalnum():
                left +=1
                continue
            elif not rightChar.isalnum():
                right -= 1
                continue
            
            if leftChar != rightChar:
                return False

            left += 1
            right -=1
        
        return True