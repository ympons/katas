# 1726. Tuple with Same Product

Given an array `nums` of **distinct** positive integers, return the number of tuples `(a, b, c, d)` such that `a * b = c * d` where `a`, `b`, `c`, and `d` are elements of `nums`, and `a != b != c != d`.

Example:
```
Input: nums = [2,3,4,6]
Output: 8
Explanation: There are 8 valid tuples:
(2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
(3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
```

**Constraints:**
- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 10^4`
- All elements in `nums` are **distinct**.
