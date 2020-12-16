# 80. Remove Duplicates from Sorted Array II

Given a **sorted** array `nums`, remove the duplicates **in-place** such that duplicates appeared at most twice and return the new length.

**Do not allocate extra space** for another array; you must do this by **modifying the input array in-place** with `O(1)` extra memory.

**Example 1:**
```
Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3]
Explanation: Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively. It doesn't matter what you leave beyond the returned length.
```

**Constraints:**
- `0 <= nums.length <= 3 * 10^4`
- `10^4 <= nums[i] <= 10^4`
- `nums` is sorted in ascending order
