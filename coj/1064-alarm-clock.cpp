#include <stdio.h>
int h1, m1, h2, m2, x1, x2;

int main(){
    scanf("%d%d%d%d", &h1,&m1,&h2,&m2);
    while (h1!=0 || m1!=0 || h2!=0 || m2!=0){
        x1 = h1*60 + m1;
        x2 = h2*60 + m2;
        if (x2 < x1){ x2 += 1440; }
        printf("%d\n", x2-x1);
        scanf("%d%d%d%d", &h1,&m1,&h2,&m2);
    }
    return 0;
}
