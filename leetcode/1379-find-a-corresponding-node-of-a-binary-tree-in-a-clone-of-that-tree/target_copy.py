# TreeNode represents a binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        def preorder(o: TreeNode, c: TreeNode):
            if o:
                if o is target:
                    self.node = c
                preorder(o.left, c.left)
                preorder(o.right, c.right)

        preorder(original, cloned)
        return self.node
