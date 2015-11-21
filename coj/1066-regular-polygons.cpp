#include <stdio.h>
#include <math.h>
#define PI 3.141592653589793

double r, pi, pc; int n;

int main(){
    int t; double angulo;
    scanf("%d", &t);
    while (t--){
        scanf("%lf %d", &r, &n);
        angulo = ((n-2)*PI/n)/2;
        pi = 2*r*n*cos(angulo);
        pc = 2*r*n/tan(angulo);
        printf("%.4lf %.4lf\n",pi,pc);
    }
    return 0;
}
