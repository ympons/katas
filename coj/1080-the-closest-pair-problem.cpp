#include <stdio.h>
#include <math.h>
struct Punto {double x, y;}puntos[10001];
int n, cantP;
double dist;

int main() {
    scanf("%d", &n);
    while (n--) {
        dist = 100000000; cantP = 0;
        scanf("%lf %lf", &puntos[cantP++].x, &puntos[cantP++].y);
        while (n--) {
            scanf("%lf %lf", &puntos[cantP].x, &puntos[cantP].y);
            double xx , yy;
            for (int k = 0; k < cantP; k++) {
                xx = (puntos[k].x - puntos[cantP].x)*(puntos[k].x - puntos[cantP].x);
                yy = (puntos[k].y - puntos[cantP].y)*(puntos[k].y - puntos[cantP].y);
                if ((xx <= dist && yy <= dist) && (xx < 100000000 && yy < 100000000) && (xx + yy < dist)) dist = xx+yy;
            }
            cantP++;
        }
        if (dist < 100000000)
            printf("%.4lf\n", sqrt(dist));
        else
            printf("INFINITY \n");
        scanf("%d", &n);
    }
    return 0;
}