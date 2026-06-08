from typing import List

# Design an algorithm to encode a list of strings to a single string. 
# The encoded string is then decoded back to the original list of strings.
# It is up to me to decide how I will decode and encode the string
class Solution:
    
    def __init__(self):
        # Separation characters
        self.START = chr(2) # Start of the string
        self.END = chr(3) # End of the string
    
    def encode(self, input: List[str]) -> str:
        
        output = ''
        
        for str in input:
            output += (self.START + str + self.END)
        
        return output
    
    def decode(self, input: str) -> List[str]:
        
        decoded = []
        
        for index, char in enumerate(input):
            if char == self.START:
                start_pos = index + 1
            if char == self.END:
                word = input[start_pos:index]
                
                decoded.append(word)
                start_pos = None
        
        print(decoded)
        return decoded
