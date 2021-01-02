# 1379. Find a Corresponding Node of a Binary Tree in a Clone of That Tree

Given two binary trees `original` and `cloned` and given a reference to a node `target` in the original tree.

The `cloned` tree is a copy of the `original` tree.

*Return a reference to the same node in the `cloned` tree.*

**Note** that you are **not allowed** to change any of the two trees or the `target` node and the answer must be a reference to a node in the `cloned` tree.

**Example:**
```
Input: tree = [7,4,3,null,null,6,19], target = 3
Output: 3
Explanation: In all examples the original and cloned trees are shown. The target node is a green node from the original tree. The answer is the yellow node from the cloned tree.
```

**Constraints:**
- The number of nodes in the `tree` is in the range `[1, 10^4]`.
- The values of the nodes of the `tree` are **unique**.
- The `target` node is a node from the `original` tree and is **not null**.
