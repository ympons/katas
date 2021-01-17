#include <stdio.h>
#define ll long long

int main() {
	ll n;
	scanf("%lld", &n);

	ll curr = 0, sum = n*(n+1)/2;
	int a; n--;
	while (n--) {
		scanf("%d", &a);
		curr += a;
	}

	printf("%lld", sum - curr);
	return 0;
}
