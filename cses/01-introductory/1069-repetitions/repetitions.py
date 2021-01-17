def max_substring_size(str):
	n = len(str)
	size = curr = 1
	for i in range(n):
		if i + 1 < n and str[i] == str[i+1]:
			curr += 1
			if size < curr:
				size = curr 
		else:
			curr = 1
	return size

if __name__ == '__main__':
	str = input().strip()
	print(max_substring_size(str))
