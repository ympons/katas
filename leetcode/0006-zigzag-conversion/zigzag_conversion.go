package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	r := make([]byte, 0, n)

	maxStep := (numRows - 1) * 2
	step1, step2 := maxStep, 0

	for i := 0; i < numRows && i < n; i++ {
		step2 = i * 2
		for j, k := i, 0; j < n; k++ {
			r = append(r, s[j])
			if step1 != maxStep && k&1 == 1 {
				j += step2
			} else {
				j += step1
			}
		}
		step1 -= 2
		if step1 < 2 {
			step1 = maxStep
		}
	}
	return string(r)
}

//
// P A Y P A L I S H I R I N G
// numRows = 1
//
//
// P Y A I H R N
// A P L S I I G
// numRows = 2
// r = PYAIHRNAPLSIIG
//
//
// P   A   H   N
// A P L S I I G
// Y   I   R
//
// P  A  Y  P  A  L  I  S  H  I  R   I   N   G
// 0  1  2  3  4  5  6  7  8  9  10  11  12  13
// 0           1           2             3        0  +4
//    4     5     6     7     8      9       10   1  +2
//       11          12          13               2  +4
// numRows = 3
//
//
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
//
// P  A  Y  P  A  L  I  S  H  I  R   I   N   G
// 0  1  2  3  4  5  6  7  8  9  10  11  12  13
// 0                 1                   2          0  +6
//    3           4     5            6       7      1  +4  +2  +4  +2
//       8     9           10    11                 2  +2  +4  +2  +4
//          12                13                    3  +6
// numRows = 4
// r = PINALSIGYAHRPI
//
//
// P     H
// A   S I
// Y  I  R
// P L   I G
// A     N
//
// P  A  Y  P  A  L  I  S  H  I  R   I   N   G
// 0  1  2  3  4  5  6  7  8  9  10  11  12  13
// 0                       1                        0  +8
//    2                 3     4                     1  +6  +2  +6  +2
//       5           6           7                  2  +4  +4
//          8     9                  10      11     3  +2  +6  +2  +6
//             12                        13         4  +8
// numRows = 5
// r = PHASIYIRPLIGAN
//
//
// P      R
// A    I I
// Y   H  N
// P  S   G
// A I
// L
//
// P  A  Y  P  A  L  I  S  H  I  R   I   N   G
// 0  1  2  3  4  5  6  7  8  9  10  11  12  13
// 0                              1                 0  +10
//    2                       3      4              1  +8   +2  +8
//       5                 6             7          2  +6   +4  +6
//          8           9                    10     3  +4   +6  +4
//             11    12                             4  +2   +8  +2
//                13                                5  +10
// numRows = 6
// r = PRAIIYHNPSGAIL
