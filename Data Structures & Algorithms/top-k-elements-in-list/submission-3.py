class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        frequency = {}

        for i in nums:
            if i in frequency:
                frequency[i] += 1
            else:
                frequency[i] = 1
        
        sorted_frequency = dict(sorted(frequency.items(), key=lambda item: item[1], reverse=True))
        first_k_numbers = list(sorted_frequency.keys())[:k]

        return first_k_numbers
