# Definition for a pair.
# class Pair:
#     def __init__(self, key: int, value: str):
#         self.key = key
#         self.value = value
class Solution:
    def insertionSort(self, pairs: List[Pair]) -> List[List[Pair]]:
        size = len(pairs)
        output = []

        for i in range(size):
            j = i - 1

            while j >= 0 and pairs[j].key > pairs[j + 1].key:
                # swap
                pairs[j], pairs[j + 1] = pairs[j + 1], pairs[j]
                
                j = j - 1
            
            output.append(pairs[:])

        return output
        