package leetcode

func maximumTime(time string) string {
	b := []byte(time)
	if b[0] == '?' {
		if b[1] == '?' || b[1] <= '3' {
			b[0] = '2'
		} else {
			b[0] = '1'
		}
	}
	if b[1] == '?' {
		if b[0] == '?' || b[0] == '2' {
			b[1] = '3'
		} else {
			b[1] = '9'
		}
	}
	if b[3] == '?' {
		b[3] = '5'
	}
	if b[4] == '?' {
		b[4] = '9'
	}
	return string(b)
}
