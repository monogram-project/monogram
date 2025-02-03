# AVL tree implementation in Python

import sys

# Adapted from https://www.programiz.com/dsa/avl-tree.

class AVLTree:

    def __init__(self, default=None):
        self.root = None
        self.default = default

    def __getitem__(self, key):
        return get_value(self.root, key, self.default)
    
    def __setitem__(self, key, value):
        if value is self.default:
            self.root = delete_node(self.root, key)
        else:
            self.root = update_or_insert_node(self.root, key, value)

    def update(self, key, value):
        if value is self.default:
            self.root = delete_node(self.root, key)
        else:
            self.root = update_or_insert_node(self.root, key, value)

    def delete(self, key):
        self.root = delete_node(self.root, key)

    def __iter__(self):
        if self.root:
            yield from self.root

    def __bool__(self):
        return self.root is not None

# Create a tree node
class TreeNode:

    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.left = None
        self.right = None
        self.height = 1

def get_value(root, key, default):
    if not root:
        return default
    if key == root.key:
        return root.value
    if key < root.key:
        return get_value(root.left, key, default)
    return get_value(root.right, key, default)

def getHeight(t):
    if not t:
        return 0
    return t.height

def getBalance(self):
    if not self:
        return 0
    return getHeight(self.left) - getHeight(self.right)

def reviseHeight(self):
    if self:
        self.height = 1 + max(getHeight(self.left), getHeight(self.right))

def getMinValueNode(self):
    root = self
    while root and root.left:
        root = root.left
    return root

def __iter__(self):
    if self:
        if self.left:
            yield from self.left
        yield self.key
        if self.right:
            yield from self.right

# Function to perform left rotation.
def leftRotate(self):
    y = self.right
    T2 = y.left
    y.left = self
    self.right = T2
    reviseHeight(self)
    reviseHeight(y)
    return y

# Function to perform right rotation.
def rightRotate(self):
    y = self.left
    T3 = y.right
    y.right = self
    self.left = T3
    reviseHeight(self)
    reviseHeight(y)
    return y

def update_or_insert_node(root, key, value):
    # Find the correct location and insert the node
    if not root:
        return TreeNode(key, value)
    elif key == root.key:
        root.value = value
        return root
    elif key < root.key:
        root.left = update_or_insert_node(root.left, key, value)
    else:
        root.right = update_or_insert_node(root.right, key, value)

    # Update the balance factor.
    reviseHeight(root)
    balanceFactor = getBalance(root)

    # Rebalance the tree.
    if balanceFactor > 1:
        # The left side of the tree is heavier - and must be truthy.
        if key < root.left.key:
            return rightRotate(root)
        else:
            root.left = leftRotate(root.left)
            return rightRotate(root)

    if balanceFactor < -1:
        # The right side of the tree is heavier - and must be truthy.
        if key > root.right.key:
            return leftRotate(root)
        else:
            root.right = rightRotate(root.right)
            return leftRotate(root)

    return root

# Function to delete a node
def delete_node(root, key):
    # Find the node to be deleted and remove it
    if not root:
        return root
    elif key < root.key:
        root.left = delete_node(root.left, key)
    elif key > root.key:
        root.right = delete_node(root.right, key)
    else:
        if not root.left:
            temp = root.right
            root = None
            return temp
        elif not root.right:
            temp = root.left
            root = None
            return temp
        temp = getMinValueNode(root.right)
        root.key = temp.key
        root.right = delete_node(root.right, temp.key)

    if not root:
        raise Exception('Cannot happen')

    # Update the balance factor of nodes
    reviseHeight(root)

    balanceFactor = getBalance(root)

    # Balance the tree
    if balanceFactor > 1:
        if getBalance(root.left) >= 0:
            return rightRotate(root)
        else:
            root.left = leftRotate(root.left)
            return rightRotate(root)
    if balanceFactor < -1:
        if getBalance(root.right) <= 0:
            return leftRotate(root)
        else:
            root.right = rightRotate(root.right)
            return leftRotate(root)
    return root


if __name__ == '__main__':
    import random
    for i in range(1000):
        avl = AVLTree()
        nums = [ i for i in range(0, 100) ]
        random.shuffle( nums )
        before = nums.copy()
        #print('nums', nums)
        for num in nums:
            avl[num] = str(num + 1000)
        random.shuffle( nums )
        after = nums.copy()
        print('50:', avl[50])
        for num in nums:
            avl.delete(num)
        if avl:
            print('BROKEN')
            break
