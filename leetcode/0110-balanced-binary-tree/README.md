# 110. Balanced Binary Tree

Given **a binary tree**, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

```
a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
```

Example 1:
```
Input: root = []
Output: true
```

Example 2:
```
Input: root = [3,9,20,null,null,15,7]
Output: true
```

Example 3:
```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
```

**Constraints:**
- The **number of nodes** in the tree is in the range `[0, 5000]`
- `-10^4 <= Node.val <= 10^4`
