# Definition for a pair.
# class Pair:
#     def __init__(self, key: int, value: str):
#         self.key = key
#         self.value = value
class Solution:
    def insertionSort(self, pairs: List[Pair]) -> List[List[Pair]]:
        output = []
        
        for i in range(0,  len(pairs)):
            print(f"iteration: {i}")

            j = i - 1
            print(f"j = {j}")
            while j >= 0 and pairs[j].key > pairs[j + 1].key:
                print(f"swap, iter:{i}")
                temp = pairs[j]
                pairs[j] = pairs[j + 1]
                pairs[j + 1] = temp
            
                j = j - 1

            output.append(pairs[:])
        
        return output

        