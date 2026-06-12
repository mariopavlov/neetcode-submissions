class Solution:
    def isValid(self, s: str) -> bool:

        parentheses = {
            ")": "(",
            "}": "{",
            "]": "["
        }

        input = []
        
        for char  in s:
            if char in parentheses:
                if len(input) == 0:
                    return False

                matching_bracket = parentheses[char]
                bracket = input.pop()

                if bracket == matching_bracket:
                    continue
                else:
                    return False
            
            input.append(char)

        return len(input) == 0
        