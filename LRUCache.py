class Node:

    def __init__(self, value, prev=None, nxt=None):
        self.value = value
        self.prev = prev
        self.nxt = nxt

class Link:

    def __init__(self):
        self.head = Node(-1)
        self.ln = 0

    def append(self, node: Node) -> int:
        hd = self.head
        ln = self.ln
        self.ln += 1
        if ln == 0:
            self.head.nxt = node
            node.prev = self.head
            self.head.prev = node
            node.nxt = self.head
            return self.ln
        else:
            self.head.prev.nxt = node
            node.prev = self.head.prev
            self.head.prev = node
            node.nxt = self.head
            return self.ln

    def remove(self, node: Node) -> int:
        node.prev.nxt = node.nxt
        node.nxt.prev = node.prev
        self.ln -= 1
        return self.ln


class LRUCache:

    def __init__(self, capacity: int):
        self.cap = cap
        self.ln = 0
        self.link = Link()
        self.item = dict()

    def get(self, key: int) -> int:
        if key in self.item:
            node = self.item[key]
            self.link.remove(node)
            self.link.append(node)
            return node.value

        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.item:
            node = self.item[key]
            self.link.remove(node)
            self.link.append(node)
            node.value = value
            return

        if self.ln < self.cap:
            node = Node(value=value)
            self.link.append(node)
            self.item[key] = node
            self.ln += 1
            return

        node = self.link.head.nxt
        self.link.remove(node)
        node = Node(value=value)
        self.link.append(node)
        self.item[key] = node
        return



# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
