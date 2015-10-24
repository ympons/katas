#include <math.h>
#include <stdio.h>

long int n;

int Potencia() {
    int raiz = (int)(sqrt(fabs((double)n))+1);
    for(int i = 2; i <= raiz; i++) {
        int p = 0, aux = n;
        while(aux%i==0) {
            p++; aux/=i;
        }
        if(fabs(aux)==1)
           if ((n>0) || (n < 0 && p & 1 != 0)) return p;
    }
    return 1;
}

int main() {
    scanf("%ld", &n);
    while(n!=0) {
        printf("%d\n", Potencia());
        scanf("%ld", &n);
    }
    return 0;
}
