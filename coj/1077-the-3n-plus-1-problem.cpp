#include <stdio.h>
#define SIZE 1000001

static unsigned short table[SIZE];

unsigned short calcCycleLength( register unsigned long n ) {
	if (n < SIZE && table[n]) return table[n];
	if (n & 1) { /* if n is odd */
		if (n < SIZE) {
			table[n] = 2 + calcCycleLength( (3*n+1) >> 1 ); /* calc two steps at one */
			return table[n];
		}
		return 2 + calcCycleLength( (3*n+1) >> 1 );
	} else {
		if( n < SIZE) {
			table[n] = 1 + calcCycleLength( n >> 1 ); /* n/2 */
			return table[n];
		}
		return 1 + calcCycleLength( n >> 1 );
	}
}

int main() {
	unsigned long i, fstIn, sndIn;
	short out = 0, temp;

	table[1] = 1;
	while ( scanf("%lu %lu", &fstIn, &sndIn ) != EOF  ) {
		if (fstIn > sndIn) {
			for( i = sndIn; i <= fstIn; ++i ) {
				temp = calcCycleLength(i);
				if (temp > out) out = temp;
			}
		} else {
			for (i = fstIn; i <= sndIn; ++i) {
				temp = calcCycleLength(i);
				if( temp > out) out = temp;
			}
		}
		printf("%lu %lu %hd\n", fstIn, sndIn, out);
		out = 0;
	}
	return 0;
}