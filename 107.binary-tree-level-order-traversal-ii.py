#
# @lc app=leetcode id=107 lang=python3
#
# [107] Binary Tree Level Order Traversal II
#
# https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
#
# algorithms
# Easy (46.27%)
# Total Accepted:    218.8K
# Total Submissions: 472.8K
# Testcase Example:  '[3,9,20,null,null,15,7]'
#
# Given a binary tree, return the bottom-up level order traversal of its nodes'
# values. (ie, from left to right, level by level from leaf to root).
# 
# 
# For example:
# Given binary tree [3,9,20,null,null,15,7],
# 
# ⁠   3
# ⁠  / \
# ⁠ 9  20
# ⁠   /  \
# ⁠  15   7
# 
# 
# 
# return its bottom-up level order traversal as:
# 
# [
# ⁠ [15,7],
# ⁠ [9,20],
# ⁠ [3]
# ]
# 
# 
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from typing import List
from collections import deque

class Solution:
    def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
        if root is None:
            return []

        stack = deque()
        result = []  # type: List[List[int]]

        result.append([root.val])
        if root.left:
            stack.append(root.left)
        if root.right:
            stack.append(root.right)

        while len(stack) != 0:
            data = []
            node = []
            while len(stack) != 0:
                item = stack.popleft()
                data.append(item.val)
                if item.left:
                    node.append(item.left)
                if item.right:
                    node.append(item.right)
            result.append(data)
            for i in node:
                stack.append(i)

        result.reverse()
        return result

                