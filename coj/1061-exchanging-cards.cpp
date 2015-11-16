#include <stdio.h>

int a, b, mark, nomark,na,nb;

int main() {
    int valor; scanf("%d%d", &a, &b);
    while (a!=0 || b!= 0) {
        bool mka[100000] = {0}, mkb[100000] = {0};
        na = a;  nb = b;
        for (int i = 0; i < a; i++) {
            scanf("%d", &valor);
            if (mka[valor]) na--; else mka[valor] = true;
        }
        mark = 0; nomark = 0;
        for (int i = 0; i < b; i++) {
            scanf("%d", &valor);
            if (mkb[valor])
                nb--;
            else {
                mkb[valor] = true;
                if (mka[valor]) mark++; else nomark++;
            }
        }
        if (na > nb) printf("%d\n", nomark); else printf("%d\n", na-mark);
        scanf("%d%d", &a, &b);
    }
    return 0;
}
