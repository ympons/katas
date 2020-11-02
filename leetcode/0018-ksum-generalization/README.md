# Ksum

Given an array `nums` of `n` integers and an integer `target`, are there `k` elements in `nums` such that `a1 + a2 + ... + ak = target`? Find all the unique sets of `k` elements in the array which gives the sum of `target`.

**Notice** that the solution must not contain duplicate sets.

**Example 1:**
```
Input: nums = [1,0,-1,0,-2,2], target = 0, k = 4
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
```

**Example 2:**
```
Input: nums = [], target = 0, k = 5
Output: []
```

**Example 3:**
```
Input: nums = [-1,0,1,2,-1,-4], target = 0, k = 3
Output: [[-1,-1,2],[-1,0,1]]
```

**Constraints:**
- `3 <= k <= 6`
- `0 <= nums.length <= 200`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
