#include <stdio.h>
int b, n, ri[100];

int main(){
    int d, c , v; bool mk;
    scanf("%d%d",&b,&n);
    while (b!=0 || n!=0){
        for (int i = 1; i <= b; i++) scanf("%d",&ri[i]);
        while (n--){
            scanf("%d%d%d",&d,&c,&v);
            ri[c] += v; ri[d] -= v;
        }
        mk = true;
        for (int i = 1; i <=b; i++)
           if (ri[i] < 0) { mk = false; break; }
        if (mk) printf("S\n"); else printf("N\n");
        scanf("%d%d",&b,&n);
    }
    return 0;
}
