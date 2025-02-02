# AVL tree implementation in Python

from abc import ABC, abstractmethod
import sys

# Adapted from https://www.programiz.com/dsa/avl-tree.

class Node(ABC):

    @abstractmethod
    def __iter__(self): ...

    @abstractmethod
    def __bool__(self): ...

    @abstractmethod
    def getBalance(self): ...

    @abstractmethod
    def reviseHeight(self): ...

    @abstractmethod
    def delete_node(self, key): ...

    @abstractmethod
    def insert_node(self, key): ...
    
    def _printHelper(self, indent, last):
        if self:
            sys.stdout.write(indent)
            p = "+----"
            i = "     "
            if last == 'R':
                p = "R----"
                i = "     "
            elif last == 'L':
                p = "L----"
                i = "|    "
            sys.stdout.write(p)
            indent += i
            print(self.key)
            self.left._printHelper(indent, 'L')
            self.right._printHelper(indent, 'R')

    # Print the tree
    def print(self, indent="", last=None):
        self._printHelper(indent, last)
  

class NullNode(Node):

    @property
    def key(self):
        return self
    
    @property
    def left(self):
        return self
    
    @property
    def right(self):
        return self
    
    @property
    def height(self):
        return 0
    
    def __bool__(self):
        return False
    
    def getBalance(self):
        return 0
    
    def reviseHeight(self):
        pass
    
    def __iter__(self):
        yield from ()

    def delete_node(self, key):
        return self
    
    def insert_node(self, key):
        return TreeNode(key)
    
    def as_str(self):
        return '.'
    

NULLNODE = NullNode()


# Create a tree node
class TreeNode(Node):

    def __init__(self, key):
        self.key = key
        self.left = NULLNODE
        self.right = NULLNODE
        self.height = 1

    def __bool__(self):
        return True
    
    def as_str(self):
        return f"({self.left.as_str()}<{self.key}<{self.right.as_str()})"
    
    def getBalance(self):
        return self.left.height - self.right.height

    def reviseHeight(self):
        self.height = 1 + max(self.left.height, self.right.height)
    
    def getMinValueNode(self):
        root = self
        while root and root.left:
            root = root.left
        return root
    
    def __iter__(self):
        yield from self.left
        yield self.key
        yield from self.right

    # Function to perform left rotation.
    def leftRotate(self):
        y = self.right
        T2 = y.left
        y.left = self
        self.right = T2
        self.reviseHeight()
        y.reviseHeight()
        return y

    # Function to perform right rotation.
    def rightRotate(self):
        y = self.left
        T3 = y.right
        y.right = self
        self.left = T3
        self.reviseHeight()
        y.reviseHeight()
        return y

    
    # Function to insert a node
    def insert_node(root, key):
        # Find the correct location and insert the node
        if key < root.key:
            root.left = root.left.insert_node(key)
        else:
            root.right = root.right.insert_node(key)

        # Update the balance factor.
        root.reviseHeight()
        balanceFactor = root.getBalance()

        # Rebalance the tree.
        if balanceFactor > 1:
            # The left side of the tree is heavier - and must be truthy.
            if key < root.left.key:
                return root.rightRotate()
            else:
                root.left = root.left.leftRotate()
                return root.rightRotate()

        if balanceFactor < -1:
            # The right side of the tree is heavier - and must be truthy.
            if key > root.right.key:
                return root.leftRotate()
            else:
                root.right = root.right.rightRotate()
                return root.leftRotate()

        return root

    # Function to delete a node
    def delete_node(root, key):
        if not root:
            Exception('Cannot happen')

        # Find the node to be deleted and remove it
        if key < root.key:
            root.left = root.left.delete_node(key)
        elif key > root.key:
            root.right = root.right.delete_node(key)
        else:
            if not root.left:
                temp = root.right
                root = None
                return temp
            elif not root.right:
                temp = root.left
                root = None
                return temp
            temp = root.right.getMinValueNode()
            root.key = temp.key
            root.right = root.right.delete_node(temp.key)

        if not root:
            raise Exception('Cannot happen')

        # Update the balance factor of nodes
        root.reviseHeight()

        balanceFactor = root.getBalance()

        # Balance the tree
        if balanceFactor > 1:
            if root.left.getBalance() >= 0:
                return root.rightRotate()
            else:
                root.left = root.left.leftRotate()
                return root.rightRotate()
        if balanceFactor < -1:
            if root.right.getBalance() <= 0:
                return root.leftRotate()
            else:
                root.right = root.right.rightRotate()
                return root.leftRotate()
        return root


if __name__ == '__main__':
    import random
    while True:
        root = NULLNODE
        nums = [ i for i in range(0, 100) ]
        random.shuffle( nums )
        before = nums.copy()
        #print('nums', nums)
        for num in nums:
            root = root.insert_node(num)
        random.shuffle( nums )
        after = nums.copy()
        #print('nums', nums)
        for num in nums:
            root = root.delete_node(num)
        if root != NULLNODE:
            print('root', root.as_str())
            print('before', before)
            print('after', after)
            print('root', list(root))
            break



