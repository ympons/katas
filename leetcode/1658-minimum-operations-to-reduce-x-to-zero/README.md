# 1658. Minimum Operations to Reduce X to Zero

You are given an integer array `nums` and an integer `x`. In one operation, you can either remove the `leftmost` or the `rightmost` element from the array `nums` and subtract its value from `x`. Note that this modifies the array for future operations.

*Return **the minimum number** of operations to reduce `x` to exactly `0` if it's possible, otherwise, return `-1`.*


**Example:**
```
Input: nums = [3,2,20,1,1,3], x = 10
Output: 5
Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.
```

**Constraints:**
- 1 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^4
- 1 <= x <= 10^9
