#include <stdio.h>
long n, cant = 0;
int cad[1000], k;

void Binario(long x) {
  k = 0;
  cant = 0; 
  while (x!=0){
    cad[k++] = x%2;
    cant += x%2;
    x /=2;
  }
}

int main() {
  scanf("%ld", &n);
  while (n!=0) {
    Binario(n);
    printf("The parity of ");
    for (int i = k-1; i > -1; i--)
      printf("%d", cad[i]);
    printf(" is %ld (mod 2).\n", cant);
    scanf("%ld", &n);
  }
  return 0;
}
