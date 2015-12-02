#include <stdio.h>
#include <stdlib.h>

struct Team {char cad[21]; int pos;};
Team teams[10001];
int n, norank[10001], lon;

long long int badness;

void QSort(int izq, int der) {
    int i = izq, j = der, v = teams[norank[(i+j)/2]].pos, temp;
    do {
        while (teams[norank[i]].pos < v) i++;
        while (teams[norank[j]].pos > v) j--;
        if (i<=j) {
            temp = norank[i]; norank[i] = norank[j]; norank[j] = temp;
            i++; j--;
        }
    }
    while (i<=j);
    if (izq < j) QSort(izq, j);
    if (der > i) QSort(i, der);
}

int main() {
    int t; scanf("%d", &t);
    while (t--) {
        int rank[10002] = {0};  lon = 0;
        scanf("%d", &n);
        for (int i = 1; i <= n; i++) {
            scanf("%s %d", teams[i].cad, &teams[i].pos);
            if (!rank[teams[i].pos])
                rank[teams[i].pos] = i;
            else
                norank[++lon] = i;
        }
        QSort(1,lon);
        badness = 0;
        for (int i = 1, j = 1; (i <= n && j <= lon); i++)
            if (!rank[i]) {
                badness += abs(teams[norank[j++]].pos - i);
            }
        printf("%lld\n", badness);
    }
    return 0;
}
