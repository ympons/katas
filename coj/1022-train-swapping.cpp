#include <stdio.h>

int lista[52], aux[52], l, solve, swap;

void Seleccion() {   
  int temp; swap = 0;
  for (int i = 0; i < l-1; i++)
    for (int j = i+1; j < l; j++)
      if (lista[i] > lista[j]) {
              temp = lista[i]; lista[i] = lista[j]; lista[j] = temp; swap++;
      }
}

void Buble() {   
  bool cambio;
  int temp; swap = 0;
  do {
    cambio = false;
    for (int i = 0; i < l-1; i++)
        if (lista[i] > lista[i+1]) {
            temp = lista[i]; lista[i] = lista[i+1]; lista[i+1] = temp; swap++;
            cambio = true;
        }
  } while (cambio);
}

int main() {
    int t; scanf("%d", &t);
    while (t--) {
        scanf("%d", &l);
        for (int i = 0; i < l; i++) {
            scanf("%d", &lista[i]);
            aux[i] = lista[i];
        }
        solve = l*(l-1);
        Seleccion(); if (swap < solve) solve = swap;
        for (int i = 0; i < l; i++) lista[i] = aux[i];
        Buble();     if (swap < solve) solve = swap;
        printf("Optimal train swapping takes %d swaps.\n", solve);
    }
    return 0;
}
