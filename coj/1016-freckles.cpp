#include <stdio.h>
#include <math.h>
struct Arista{ int ini, fin; double cost;};
struct Punto{ double x, y;};

Arista aristas[10001]; Punto puntos[101];
int k, n; double minimo;

double Distancia(double x1, double y1, double x2, double y2){return sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2));}

void QSort(int izq, int der) {
	int i = izq, j = der;
	double v = aristas[(i+j)/2].cost;
	do{
		while (aristas[i].cost < v) i++;
		while (aristas[j].cost > v) j--;
		if (i<=j) {
			Arista temp = aristas[i];
			aristas[i] = aristas[j];
			aristas[j] = temp;
			i++; j--;
		}
	}while(i<=j);
	if (izq < j) QSort(izq, j);
	if (der > i) QSort(i, der);
}

void Kruskal() {
	int j, c = 0, x, y, Comp[10001];
	for (int i = 1; i <= k; i++)
		if (Comp[aristas[i].ini]==0 && Comp[aristas[i].fin]==0)	{
			c++; Comp[aristas[i].ini] = c; Comp[aristas[i].fin] = c;
		}
		else if (Comp[aristas[i].ini]*Comp[aristas[i].fin] == 0) {
			if (Comp[aristas[i].ini] == 0)
				j = aristas[i].ini;
			else
				j = aristas[i].fin;
			Comp[j] = Comp[aristas[i].ini] + Comp[aristas[i].fin];
		} else if (Comp[aristas[i].ini] != Comp[aristas[i].fin]) {
			x = Comp[aristas[i].ini];  y = Comp[aristas[i].fin];
			for (int j = 1; j <= n; j++)
				if (Comp[j] == x) Comp[j] = y;
		}
		else aristas[i].cost = 0.0;
		minimo = 0.0;
		for (int i = 1; i <= k; i++)
			if (aristas[i].cost != 0.0) minimo += aristas[i].cost;
}

int main() {
    scanf("%d", &n);
    for (int i = 1; i <= n; i++)
        scanf("%lf %lf", &puntos[i].x, &puntos[i].y);
    k = 0;
    for (int i = 1; i < n; i++)
        for (int j = i+1; j <= n; j++) {
            aristas[++k].ini  = i;
            aristas[k].fin  = j;
            aristas[k].cost = Distancia(puntos[i].x, puntos[i].y, puntos[j].x, puntos[j].y);
        }
    QSort(1, k);
    Kruskal();
    printf("%.2lf", minimo);
    return 0;
}
