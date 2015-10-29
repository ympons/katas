#include <stdio.h>
int main() {
 float sum = 0, x;
 for (int i=0; i<12; i++) {
  scanf("%f",&x);
  sum += x;
 };
 printf("$%0.2f",sum/12);
 return 0;
};
