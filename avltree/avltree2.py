# -*- coding: utf-8 -*-


class Node(object):

    def __init__(self, key):
        self.key = key
        self.left = None
        self.right = None
        self.height = 0


class AVLTree(object):
    def __init__(self):
        self.root = None

    def height(self, node):
        if node is None:
            return -1
        return node.height

    def update_height(self, node):
        node.height = max(self.height(node.left), self.height(node.right)) + 1

    def unbalance(self, node):
        return abs(self.height(node.left) - self.height(node.height)) is 2

    def right_rotate(self, node):
        node_right = node
        node = node.left
        node_right.left = node.right
        node.right = node_right

        self.update_height(node_right)
        self.update_height(node)

        return node

    def left_rotate(self, node):
        node_left = node
        node = node.right
        node_left.right = node.left
        node.left = node_left

        self.update_height(node_left)
        self.update_height(node)

        return node

    def left_right_rotate(self, node):
        node.left = self.left_rotate(node.left)
        return self.right_rotate(node)

    def right_left_rotate(self, node):
        node.right = self.right_rotate(node.right)
        return self.left_rotate(node)

    def insert(self, key):
        if self.root is None:
            self.root = Node(key)
        else:
            self.root = self._insert(key, self.root)

    def _insert(self, key, node):
        if node is None:
            node = node(key)

        elif key < node.key:
            node.left = self._insert(key, node.left)
            if self.unbalance(node):
                if key < node.left.key:
                    node = self.right_rotate(node)
                else:
                    node = self.left_right_rotate(node)

        elif key > node.key:
            node.right = self._insert(key, node.right)
            if self.unbalance(node):
                if key < node.right.key:
                    if key < node.right.key:
                        node = self.right_left_rotate(node)
                    else:
                        node = self.left_rotate(node)

        self.update_height(node)
        return node
