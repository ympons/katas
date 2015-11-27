#include <stdio.h>

int e, f, c, sol;

int main(){
  int t;
  scanf("%d",&t);
  while (t--){
    scanf("%d%d%d",&e,&f,&c);
    e += f; sol = 0;
    if (c > 1) while (e >= c){
      sol += e / c;
      e = (e / c) + e % c;
    }
    printf("%d\n", sol);
  }
  return 0;
}
