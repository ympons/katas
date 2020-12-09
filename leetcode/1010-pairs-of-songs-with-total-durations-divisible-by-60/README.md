# 1010. Pairs of Songs With Total Durations Divisible by 60

You are given a list of songs where the `ith` song has a duration of `time[i]` seconds.

Return *the number of pairs of songs* for which their total duration in seconds is divisible by `60`. Formally, we want the number of indices `i`, `j` such that `i < j` with `(time[i] + time[j]) % 60 == 0`.

**Example:**
```
Input: time = [60,60,60]
Output: 3
Explanation: All three pairs have a total duration of 120, which is divisible by 60.
```

**Constraints:**
- 1 <= time.length <= 6 * 10^4
- 1 <= time[i] <= 500
