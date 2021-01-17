#include <stdio.h>
#define ll long long

int main() {
	int n;
	scanf("%d", &n);

	ll x, prev = 0, steps = 0;
	while (n--) {
		scanf("%lld", &x);
		if (prev > 0 && x < prev) {
			steps += prev - x;
		} else {
			prev = x;
		}
	}

	printf("%lld", steps);
	return 0;
}