#include <stdio.h>
int main(){
  int t,n,h,i;
  scanf("%d", &t);
  while(t--){
    scanf("%d",&n);
    h=1;
    for (i=1;i<=n;i++)
      if (i%2) h*=2; else h++;
    printf("%d\n",h);
  }
  return 0;
}
