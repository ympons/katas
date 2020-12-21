# 1463. Cherry Pickup II

Given a `rows x cols` matrix `grid` representing a field of cherries. Each cell in `grid` represents the number of cherries that you can collect.

You have two robots that can collect cherries for you, **Robot #1** is located at the top-left corner `(0,0)`, and **Robot #2** is located at the top-right corner `(0, cols-1)` of the `grid`.

Return the **maximum number of cherries collection** using both robots by following the rules below:

- From a cell `(i,j)`, robots can move to cell `(i+1, j-1)`, `(i+1, j)` or `(i+1, j+1)`.
- When any robot is passing through a cell, It picks it up all cherries, and the cell becomes an empty cell `(0)`.
- When both robots stay on the same cell, only one of them takes the cherries.
- Both robots cannot move outside of the `grid` at any moment.
- Both robots should reach the bottom row in the `grid`.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/04/29/sample_1_1802.png)
```
Input: grid = [[3,1,1],[2,5,1],[1,5,5],[2,1,1]]
Output: 24
Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
Cherries taken by Robot #1, (3 + 2 + 5 + 2) = 12.
Cherries taken by Robot #2, (1 + 5 + 5 + 1) = 12.
Total of cherries: 12 + 12 = 24.
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/04/23/sample_2_1802.png)
```
Input: grid = [[1,0,0,0,0,0,1],[2,0,0,0,0,3,0],[2,0,9,0,0,0,0],[0,3,0,5,4,0,0],[1,0,2,3,0,0,6]]
Output: 28
Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
Cherries taken by Robot #1, (1 + 9 + 5 + 2) = 17.
Cherries taken by Robot #2, (1 + 3 + 4 + 3) = 11.
Total of cherries: 17 + 11 = 28.
```

**Constraints:**
- `rows == grid.length`
- `cols == grid[i].length`
- `2 <= rows, cols <= 70`
- `0 <= grid[i][j] <= 100`

## Solution

This is an optimization problem since we need to find the max number of cherries collection. Also, there are overlapping sub-problems which can be solved through DP.

Based on the rules of how the robots move we can establish the following assumptions:
- the robots move at the same speed.
- the robots will move M times, with M the number of rows in the grid.
- the robots are always in the same row at every step.

From the latest assumption we can deduce that it is needed only 3 dimensions to address the sub-problems: `row`, `col1`, and `col2`.

### Algorithm

Define a three-dimension `dp` array where each dimension has a size of `m`, `n`, and `n` respectively.

`dp[row][col1][col2]` represents the maximum cherries we can pick if robot1 starts at `(row, col1)` and robot2 starts at `(row, col2)`.

To compute `dp[row][col1][col2]` (transition equation):

- Collect the cherry at `(row, col1)` and `(row, col2)`. Do not double count if `col1 == col2`.
- If we are not in the last row, we need to add the maximum cherries we can pick in the future.
- The maximum cherries we can pick in the future is the maximum of `dp[row+1][new_col1][new_col2]`, where `new_col1` can be `col1`, `col1+1`, or `col1-1`, and `new_col2` can be `col2`, `col2+1`, or `col2-1`.

Finally, return dp[0][0][last_column].
