from typing import List


class Solution(object):

    def rotate(self, nums: List[int], k: int) -> None:
        for i in range(k):
          tmp = nums[-1]
          n = len(nums)
          while n > 1:
            nums[n - 1] = nums[n-2]
            n -= 1
          nums[0] = tmp

    def rotate1(self, nums: List[int], k: int) -> None:
        """use once in-place move"""
        k %= len(nums)
        count = 0
        for i in range(len(nums)):
            current = i
            prev = nums[i]
            while True:
                nxt = (current + k) % len(nums)
                tmp = nums[nxt]
                nums[nxt] = prev
                prev = tmp
                current = nxt
                count += 1
                if i == current:
                    break


    def rotate2(self, nums: List[int], k: int) -> None:
        """use reverse"""
        k %= len(nums)
        self.reverse(nums, 0, len(nums) - 1)
        self.reverse(nums, 0, k - 1)
        self.reverse(nums, k, len(nums) - 1)

    def reverse(self, nums: List[int], start: int, end: int):
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1
