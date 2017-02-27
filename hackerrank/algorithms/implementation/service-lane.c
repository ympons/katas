#include <stdio.h>
int n,t,min,a,b, w[100001];
int main() {
    scanf("%d %d", &n, &t);
    int i=0;
    while (i<n){
        scanf("%d", &w[i]);i++;
    }
    while(t--){
        scanf("%d %d",&a,&b);
        min=w[a];
        i=a+1;
        while(i<=b){
            if (min > w[i]) min = w[i];
            i++;
        }
        printf("%d\n", min);
    }
    return 0;
}

