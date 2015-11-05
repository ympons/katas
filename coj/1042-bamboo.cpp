#include <stdio.h>
float num, h , b, prom;
int t;
int main() {
	prom = 0.0;
	scanf("%d", &t)	;
	for (int i = 0; i < t; i++) {
		scanf("%f%f", &h, &b);
		num = (h*h - b*b)/(2*h);
		prom += num;
		printf("%.1f\n", num);
	}
	prom /= t;
	printf("%.1f\n", prom);
	return 0;
}
