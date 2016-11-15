#include <stdio.h>
int n, ni, cni, total;

int main() {
  scanf("%d", &n);
  for (int i=0; i<n; i++) {
    scanf("%d", &cni);
    total = 0;
    for(int j=0; j<cni; j++) {
       scanf("%d", &ni);
       total += ni;
    }
    if (total%cni==0) printf("YES\n"); else printf("NO\n");
  }
  return 0;
}