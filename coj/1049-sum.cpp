#include <stdio.h>

int main() {
 int n, k;
 scanf("%d",&n);
 if (n > 0)
  n = n*(n + 1)/2;
 else {
  k = -1*n;
  k = k*(k + 1)/2;
  n = -1*(k -1);
 }
 printf("%d", n);
 return 0;
}
