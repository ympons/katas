#include <stdio.h>

bool mk[20][20];
int n, min, solve = 0;

int minimo (int a, int b){return (a < b)? a : b;}

int main()
{
    int t; char car, car1;
    scanf("%d", &t);
    while (t--)
    {
       scanf("%d", &n);
       for (int i = 0; i < 20; i++)
          for (int j = 0; j < 20; j++) mk[i][j] = false;
       solve = 0;
       scanf("%c", &car1);
       for (int i = 0; i < n; i++)
       {
          for (int j = 0; j < n; j++)
          {
              scanf("%c", &car);
              if (car=='.') mk[i][j] = true;
          }
          scanf("%c", &car1);
       }
       bool bandera = false;
       for (int i = 0; i < n; i++)
          for (int j = 0; j < n; j++)
           if (mk[i][j])
           {
                min = minimo(n-i, n-j);
                for (int k = min; k>solve; k--)
                {
                    bandera = false;
                    for (int l1 = i; l1 < i+k; l1++)
                    {
                       for (int l2 = j; l2 < j+k; l2++)
                          if (!mk[l1][l2])
                          {
                              bandera = true;
                              break;
                          }
                       if (bandera) break;
                    }
                    if (!bandera && solve < k)
                    {
                        solve = k; break;
                    }
                }
           }
       printf("%d\n", solve);
    }
    return 0;
}
