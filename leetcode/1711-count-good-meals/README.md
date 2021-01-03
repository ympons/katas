# 1711. Count Good Meals

A **good meal** is a meal that contains **exactly two different food items** with a sum of deliciousness equal to a power of two.

Given an array of integers `deliciousness` where `deliciousness[i]` is the deliciousness of the `ith` item of food, *return the number of different good meals you can make from this list modulo `10^9 + 7`*.

Note that items with different indices are considered different even if they have the same deliciousness value.

**Example:**
```
Input: deliciousness = [1,3,5,7,9]
Output: 4
Explanation: The good meals are (1,3), (1,7), (3,5) and, (7,9).
Their respective sums are 4, 8, 8, and 16, all of which are powers of 2.
```

**Constraints:**
- `1 <= deliciousness.length <= 10^5`
- `0 <= deliciousness[i] <= 2^20`
