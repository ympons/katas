# 1640. Check Array Formation Through Concatenation

You are given an array of **distinct** integers `arr` and an array of integer arrays `pieces`, where the integers in `pieces` are **distinct**. Your goal is to form `arr` by concatenating the arrays in `pieces` **in any order**. However, you are **not** allowed to reorder the integers in each array `pieces[i]`.

Return `true` if it is possible to form the array `arr` from `pieces`. Otherwise, return `false`.

**Example 1:**
```
Input: arr = [10], pieces = [[10]]
Output: true
```

**Example 2:**
```
Input: arr = [1,3,5,7], pieces = [[2,4,6,8]]
Output: false
```

**Constraints:**
- `1 <= pieces.length <= arr.length <= 100`
- `sum(pieces[i].length) == arr.length`
- `1 <= pieces[i].length <= arr.length`
- `1 <= arr[i], pieces[i][j] <= 100`
- The integers in `arr` are **distinct**
- The integers in `pieces` are **distinct** (i.e., *If we flatten pieces in a 1D array, all the integers in this array are distinct*)
