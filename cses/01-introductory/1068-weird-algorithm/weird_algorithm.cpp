#include <stdio.h>
#define ll long long

void printNumbers(ll n) {
    printf("%lld", n);
    while (n > 1) {
        n = (n & 1) ? n*3+1 : n/2;
        printf(" %lld", n);
    }
}

int main() {
    ll n;
    scanf("%lld", &n);
    printNumbers(n);
    return 0;
}
