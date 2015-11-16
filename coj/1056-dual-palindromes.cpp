#include <stdio.h>

long long int n, m, cont;

bool Palin(long long int valor, int bas) {
   long long int x = valor, k = 0;  int cad[100];
   while (x!=0) {
      cad[k++]=x%bas;  x /=bas;
   }
   while (cad[k-1]==0 && k-1 >= 0) k--;
   for (int p=0; p<=(k/2); p++)
     if (cad[p]!=cad[k-1-p]) return false;
   return true;
}

int main() {
    while (scanf("%ld", &n)!=EOF) {
        scanf("%lld", &m); m++;
        for (int i = 0; i < n; m++) {
            cont = 0;
            for (int j = 2; j <= 10; j++)
                if (Palin(m, j)) {
                    cont++;
                    if (cont==2) {
                        printf("%lld\n", m); i++; break;
                    }
                }
        }
    }
    return 0;
}
