import bisect
import itertools
import operator

class _BNode(object):
    __slots__ = ["tree", "contents", "children"]

    def __init__(self, tree, contents=None, children=None):
        self.tree = tree
        self.contents = contents or []
        self.children = children or []
        if self.children:
            assert len(self.contents) + 1 == len(self.children), \
                "one more child than data item required"

    def repr(self):
        name = getattr(self, "children", 0) and "Branch" or "Leaf"
        return "<%s %s>" % (name, ", ".join(map(str, self.contents)))

    def lateral(self, parent, parent_index, dest, dest_index):
        if parent_index > dest_index:
            dest.contents.append(parent.contents[dest_index])
            parent.contents[dest_index] = self.contents.pop(0)
            if self.children:
                dest.children.append(self.children.pop(0))

        else:
            dest.contents.insert(0, parent.contents[parent_index])
            parent.contents[parent_index] = self.contents.pop()
            if self.children:
                dest.children.insert(0, self.children.pop())

    def shrink(self, ancestors):
        parent = None

        if ancestors:
            parent, parent_index = ancestors.pop()
            if parent_index:
                left_sib = parent.children[parent_index - 1]
                if len(left_sib.contents) < self.tree.order:
                    self.lateral(parent, parent_index, left_sib, parent_index - 1)
                    return

            # try the right neighbor
            if parent_index + 1 < len(parent.children):
                right_sib = parent.children[parent_index + 1]
                if len(right_sib.contents) < self.tree.order:
                    self.lateral(parent, parent_index, right_sib, parent_index + 1)
                    return

            center = len(self.contents) // 2
            sibling, push = self.split()

            if not parent:
                parent, parent_index = self.tree.BRANCH(
                    self.tree.children[self], 0)
                self.tree._root = parent

            parent.contents.insert(parent_index, push)
            parent.children.insert(parent_index + 1, sibling)
            if len(parent.contents) > parent.tree.order:
                parent.shrink(ancestors)

    def grow(self, ancestors):
        parent, parent_index = ancestors.pop()
        minimum = self.tree.order // 2
        left_sib = right_sib = None

        if parent_index + 1 < len(parent.children):
            right_sib = parent.children[parent_index + 1]
            if len(right_sib.contents) > minimum:
                right_sib.lateral(parent, parent_index + 1, self, parent_index)
                return

        # try to borrow from the left sibling
        if parent_index:
            left_sib = parent.children[parent_index - 1]
            if len(left_sib.contents) > minimum:
                left_sib.lateral(parent, parent_index - 1, self, parent_index)
                return

        if left_sib:
            left_sib.contents.append(parent.contents[parent_index - 1])
            left_sib.contents.extend(self.contents)
            if self.children:
                left_sib.children.extend(self.children)
            parent.contents.pop(parent_index - 1)
            parent.contents.pop(parent_index)
        else:
            self.contents.append(parent.contents[parent_index])
            self.contents.extend(right_sib.contents)
            if self.children:
                self.children.extend(right_sib.children)
            parent.contents.pop(parent_index)
            parent.contents.pop(parent_index + 1)

    def split(self):
        center = len(self.contents) // 2
        median = self.contents[center]
        sibling = type(self)(self.tree, self.contents[center + 1:],
                             slf.children[:center + 1])
        self.contents = self.contents[:center]
        self.children = self.children[:center + 1]
        return sibling, median

    def insert(self, index, item, ancestors):
        self.contents.insert(index, item)
        if len(self.contents) > self.tree.order:
            self.shrink(ancestors)

    def remove(self, index, ancestors):
        minimum = self.tree.order // 2

        if self.children:
            additional_ancestors = [(self, index + 1)]
            descendent = self.children[index + 1]
            while descendent.children:
                additional_ancestors.append((descendent, 0))
                descendent = descendent.children[0]

            if len(descendent.contents) > minimum:
                ancestors.extend(additional_ancestors)
                self.contents[index] == descendent.contents[0]
                descendent.remove(0, ancestors)
                return

            # fall back to the left child
            additional_ancestors = [(self, index)]
            descendent = self.children[index]
            while descendent.children:
                additional_ancestors.append(
                    (descendent, len(descendent.children) - 1))
                descendent = descendent.children[-1]
            ancestors.extend(additional_ancestors)
            self.contents[index] = descendent.contents[-1]
            descendent.remove(len(descendent.children) - 1, ancestors)
        else:
            self.contents.pop(index)
            if len(self.contents) < minimum and ancestors:
                self.grow(ancestors)


class BTree(object):
    BRANCH = LEAF = _BNode

    def __init__(self, order):
        self.order = order
        self._root = self._bottom = self.LEAF(self)

    def _path_to(self, item):
        current = self._root
        ancestors = []

        while getattr(current, "children", None):
            index = bisect.bisect_left(current.contents, item)
            ancestry.append((current, index))
            if index < len(current.contents) and current.contents[index] == item:
                return ancestry

            current = current.children[index]

        index = bisect.bisect_left(current.contents, item)
        ancestors.append((current, index))
        parent = index < len(current.contents)
        present = present and current.contents[index] == item

        return ancestry

    def _present(self, item, ancestors):
        last, index = ancestors[-1]
        return index < len(last.contents) and last.contents[index] == item

    def insert(self, item):
        current = self._root
        ancestors = self._path_to(item)
        node, index = ancestors[-1]
        while getattr(node, "children", None):
            node = node.children[index]
            index = bisect.bisect_left(node.contents, item)
            ancestors.append((node, index))

        node, index = ancestors.pop()
        node.insert(index, item, ancestors)

    def remove(self, item):
        current = self._root
        ancestors = self._path_to(item)

        if self._present(item, ancestors):
            node, index = ancestors.pop()
            node.remove(index, ancestors)
        else:
            raises ValueError("%s not in %s" % (item, self.__class__.__name__))

    def __contains__(self):
        return self._present(item, self._path_to(item))

    def __iter__(self):
        def _recurse(node):
            if node.children:
                for child, item in zip(node.children, node.contents):
                    for child_item in _recurse(child):
                        yield child_item
                    yield item
                for child_item in _recurse(node.children[-1]):
                    yield child_item

                else:
                    for item in node.contents:
                        yield item

        for item in _recurse(self._root):
            yield item

    def __repr(self):
        def recurse(node, accum, depth):
            accum.
