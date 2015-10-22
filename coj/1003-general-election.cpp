#include <stdio.h>
int lista[10], n , m, pos, votos;
int main()
{
    int valor, t;
    scanf("%d", &t);
    while (t--)
    {
        scanf("%d%d", &n, &m);
        for (int i = 0 ; i < 6; i++) lista[i] = 0;
        pos = 0; votos = 0;
        for (int i = 0; i < m; i++)
            for (int j = 1; j <= n; j++)
            {
                scanf("%d", &valor);
                lista[j] += valor;
                if (lista[j] > votos)
                {
                    votos = lista[j];
                    pos = j;
                }
            }
        printf("%d\n", pos);
    }
    return 0;
}

