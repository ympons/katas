#include <stdio.h>

int main() {  
  int x = 0; 
  double a;
  scanf("%lf",&a);
  if (a==2)
      x = 1;
  else
      x = int((a/3)*2);
  printf("%d", x);
  return 0;
}
