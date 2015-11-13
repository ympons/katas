#include <stdio.h>
int n,a,b; long long int suma;
int main() {
  scanf("%d", &n);
  while (n--) {
    scanf("%d%d", &a, &b);
    suma = 0;
    for (long i = a; i <=b; i++) {
      suma += i*(i+1)*(i+2);
    }
    printf("%lld\n", suma%100);
  }
  return 0;
}
