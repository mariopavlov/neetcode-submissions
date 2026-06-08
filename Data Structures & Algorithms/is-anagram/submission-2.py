class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        
        # Check the size of two input strings, and return False if they are not equal
        if len(s) != len(t):
            return False

        # Check if two strings are anagrams
        # 1. Pick a data structure,
        # sort(string)
        # Time Complexity: O(n) - O(log n)
        # Space Complexity: O(1)
        sort_s = sorted(s)
        sort_t = sorted(t)

        if sort_s == sort_t:
            return True

        # Satisfy solution
        return False