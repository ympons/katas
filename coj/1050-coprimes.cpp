#include <stdio.h>

int n, cont = 1;

int MCD(int a, int b) {
 int mcd = a;
 do {
   a = b;
   b = mcd%b;
   mcd = a;
 }while(b!=0);
 return mcd;
};

int main() {
   scanf("%d", &n);
   for (int i = 2; i<n; i++)
      if (MCD(n,i)==1) cont++;
   if (cont==1)
      printf("0\n");
   else 
      printf("%d\n", cont);
   return 0;
}
