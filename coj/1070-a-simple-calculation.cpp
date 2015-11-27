#include <stdio.h>
int n;
int main(){
  while (scanf("%d",&n)!=EOF){
    double x = n*(n+1)*(2*n+1)/6;
    double y = n*n*(n+1)*(n+1)/4;
    printf("%0.lf %0.lf\n",x,y);
  }
  return 0;
}
