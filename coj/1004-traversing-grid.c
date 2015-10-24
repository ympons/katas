#include <stdio.h>
int n , m;

int main() {
    int valor, t;
    scanf("%d", &t);
    while (t--) {
        scanf("%d%d", &n, &m);
            if(n==m) {
                if(n%2==0) printf("%s\n","L");
                  else printf("%s\n","R");
            } else {
                if(n>m) {
                    if(m%2==0) printf("%s\n","U");
                      else printf("%s\n","D");
                } else {
                    if(n%2==0) printf("%s\n","L");
                      else printf("%s\n","R");
                }
         }
    }
    return 0;
}
