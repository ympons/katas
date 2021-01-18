#include <stdio.h>

const int SIZE = int(1e6) + 1;
int P[SIZE];

int permutation(int n) {
	if (n == 2 || n == 3) return -1;
	int i = 0;
	for (int j = 2; j <= n; j += 2) {
		P[i++] = j;
	}
	for (int j = 1; i < n && j <= n; j += 2) {
		P[i++] = j;
	}
	return n;
}

int main() {
	int n;
	scanf("%d", &n);

	n = permutation(n);
	if (n == -1) {
		printf("%s", "NO SOLUTION");
		return 0;
	}

	n--; printf("%d", P[n]);
	while (n--) {
		printf(" %d", P[n]);
	}

	return 0;
}