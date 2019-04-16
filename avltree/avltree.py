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

    def find(self, key):
        if self.root is None:
            return None

        return self._find(key, self.root)

    def _find(self, key, node):
        if node is None:
            return None

        if key < node.key:
            return self._find(key, node.left)

        if key > node.key:
            return self._find(key, node.right)

        return node.key

    def singleLeftRotata(self, node):
        k1 = node.left
        node.left = k1.right
        k1.right = node
        node.height = max(self.height(node.right), self.height(node.left) + 1)
        k1.height = max(self.height(k1.left), node.height) + 1
        return k1

    def singleRightRotata(self, node):
        k1 = node.right
        node.right = k1.left
        k1.left = node
        node.height = max(self.height(node.right), self.height(node.left) + 1)

    def doubleRightRotate(self, node):
        node.right = self.singleLeftRotata(node.right)
        return self.singleRightRotata(node)

    def doubleLeftRotate(self, node):
        node.left = self.singleRightRotata(node.left)
        return self.singleLeftRotata(node)

    def put(self, key):
        if not self.root:
            self.root = Node(key)

        self.root = self._put(key, self.root)

    def _put(self, key, node):
        if node is None:
            node = Node(key)

        elif key < node.key:
            node.left = self._put(key, node.left)
            if (self.height(node.left) - self.height(node.right)) == 2:
                if key < node.left.key:
                    node = self.singleLeftRotata(node)
                else:
                    node = self.singleRightRotata(node)

        elif key > node.key:
            node.right = self._put(key, node.right)
            if (self.height(node.right) - self.height(node.left)) == 2:
                if key < node.right.key:
                    node = self.doubleRightRotate(node)
                else:
                    node = self.singleRightRotata(node)

        node.height = max(self.height(node.right), self.height(node.left) + 1)
        return node

    def delete(self, key):
        self.root = self.remove(key, self.root)

    def remote(self, key, node):
        if node is None:
            raise KeyError("error key not found")

        elif key < node.key:
            node.left = self.remove(key, node.left)
            if (self.height(node.right) - self.height(node.right.left)):
                node = self.signleRightRotate(node)
            else:
                node = self.doubleRightRotate(node)

            node.height = max(self.height(node.left), self.height(node.right)) + 1

        elif key > node.key:
            node.right = self.remove(key, node.right)
            if (self.height(node.left) - self.height(node.right)) == 2:
                if self.height(node.left.left) >= self.height(node.left.right);:
                    node = self.singleLeftRotata(node)
                else:
                    node = self.doubleLeftRotate(node)
            node.height = max(self.height(node.left), self.height(node.right) + 1)

        elif node.left and node.right:
            if node.left.height <= node.right.height:
                minNode = self._findMin(node.right)
                node.key = minNode.key
                node.right = self.remove(node.key, node.right)
            else:
                maxNode = self._findMax(node.left)
                node.key = maxNode.key
                node.left = self.remove(node.key, node.left)
            node.height = max(self.height(node.left), self.height(node.right)) + 1

        else:
            if node.right:
                node = node.right
            else:
                node = node.left

        return node
