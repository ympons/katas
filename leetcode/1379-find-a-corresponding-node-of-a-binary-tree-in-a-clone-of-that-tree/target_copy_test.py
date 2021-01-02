import unittest

from collections import deque
from target_copy import Solution, TreeNode

class SolutionTest(unittest.TestCase):
    def test_target_copy_example_1(self):
        original = self.deserialize("7,4,3,null,null,6,19")
        cloned = self.deserialize("7,4,3,null,null,6,19")
        target = original.right
        expected = cloned.right
        self.assertEqual(Solution().getTargetCopy(original, cloned, target), expected)

    def test_target_copy_example_2(self):
        original = self.deserialize("8,null,6,null,5,null,4,null,3,null,2,null,1")
        cloned = self.deserialize("8,null,6,null,5,null,4,null,3,null,2,null,1")
        target = original.right.right.right
        expected = cloned.right.right.right
        self.assertEqual(Solution().getTargetCopy(original, cloned, target), expected)

    def test_target_copy_example_3(self):
        original = self.deserialize("1,2,3,4,5,6,7,8,9,10")
        cloned = self.deserialize("1,2,3,4,5,6,7,8,9,10")
        target = original.left.right
        expected = cloned.left.right
        self.assertEqual(Solution().getTargetCopy(original, cloned, target), expected)

    def test_target_copy_example_4(self):
        original = self.deserialize("1,2,null,3")
        cloned = self.deserialize("1,2,null,3")
        target = original.left
        expected = cloned.left
        self.assertEqual(Solution().getTargetCopy(original, cloned, target), expected)

    # utilities
    def serialize(self, root):
        """Serializes a binary tree to a single string"""
        flatten = []
        queue = deque([root])
        while queue:
            node = queue.pop()
            if node:
                flatten.append(str(node.val))
                queue.appendleft(node.left)
                queue.appendleft(node.right)
            else:
                flatten.append('null')
        return ','.join(flatten)

    def deserialize(self, data):
        """Deserializes the given data to a binary tree"""
        if not data or data == 'null':
            return None
        flatten = data.split(',')
        root = TreeNode(flatten[0])
        queue = deque([root])
        i = 1
        while queue:
            node = queue.pop()
            if i < len(flatten) and flatten[i] != 'null':
                node.left = TreeNode(int(flatten[i]))
                queue.appendleft(node.left)
            i += 1
            if i < len(flatten) and flatten[i] != 'null':
                node.right = TreeNode(int(flatten[i]))
                queue.appendleft(node.right)
            i += 1
        return root

if __name__ == "__main__":
    unittest.main()
