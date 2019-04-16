from typing import List


class Solution(object):
    def generate(self, num_rows: int) -> List[List[int]]:
        result = []

        for row_num in range(num_rows):
            row = [None for _ in range(row_num)]
            row[0], row[1] = 1, 1

            for j in range(1, len(row) - 1):
                row[j] = result[row_num - 1][j-1] + result[row_num - 1][j]

            result.append(row)

        return result
