# Definition for a pair.
# class Pair:
#     def __init__(self, key: int, value: str):
#         self.key = key
#         self.value = value
class Solution:
    def insertionSort(self, pairs: List[Pair]) -> List[List[Pair]]:
        output = []
        
        for i in range(0, len(pairs)):
            j = i - 1
            while j >= 0 and pairs[j].key > pairs[j + 1].key:
                print("swap")

                temp = pairs[j]
                pairs[j] = pairs[j + 1]
                pairs[j + 1] = temp

                j = j - 1
        
            output.append(pairs[:])

        return output

    def print_pairs(self, pairs: List[Pair]):
        for pair in pairs:
            print(f"{pair.key}:{pair.value}")