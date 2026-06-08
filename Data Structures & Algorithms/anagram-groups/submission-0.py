class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        temporary = {}
        
        for elem in strs:
            sorted_elem = ''.join(sorted(elem))
            if sorted_elem in temporary:
                temporary[sorted_elem].append(elem)
            else:
                temporary[sorted_elem] = [elem]

        return [value for value in temporary.values()]
        