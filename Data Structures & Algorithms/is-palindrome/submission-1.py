class Solution:
    def isPalindrome(self, s: str) -> bool:
        left = 0
        right = len(s) - 1

        while left <= right:
            left_char = s[left].lower()
            right_char = s[right].lower()

            print("Compare: " + left_char + " == " + right_char)

            if not left_char.isalnum():
                left += 1
                continue
            if not right_char.isalnum():
                right -= 1
                continue
            
            if not left_char == right_char:
                return False
            else:
                left += 1
                right -= 1
                continue
            
        return True
        
        