#include <stdio.h>
int t,n,k,s;
int main() {
    scanf("%d", &t);
    while(t--){
        scanf("%d %d",&n,&k);
        int y = 0;
        while(n--){
            scanf("%d",&s);
            if (s <= 0) y++;
        }
        if (y>=k) 
            printf("NO\n"); 
        else
            printf("YES\n");
    }
    return 0;
}

