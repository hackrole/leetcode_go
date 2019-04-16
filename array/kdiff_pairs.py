from typing import List
from collections import Counter


class Solution:
    def findPairs(self, nums: List[int], k: int) -> int:
        if k < 0:
            return 0

        if k == 0:
            a = Counter(nums)
            return len([i for i, v in a.items() if v > 1])

        nums = set(nums)
        nums2 = map(lambda x: x + 2, nums)
        return len(nums & nums2)
