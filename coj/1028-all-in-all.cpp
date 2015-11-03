#include <stdio.h>
#include <string.h>

char cad1[100001], cad2[100001];
long n, m, i = 0, j = 0;
int main() {
   while (scanf("%s %s",&cad1,&cad2)!=EOF) {
        n = strlen(cad1);
        m = strlen(cad2);
        if (n > m) {
            printf("No\n");
        } else {
            i = 0;
            for (j = 0; (j < m) && ((m-j) >= (n-i)); j++)
               if (cad1[i]==cad2[j]) i++;
            if (i>=n)
                printf("Yes\n");
            else
                printf("No\n");
        }
   }
   return 0;
}

