#include <stdio.h>
#include <string.h>

char cad[101], solve[110]; int lon;

int main(){
  while (scanf("%s", cad)!=EOF){
    int acarreo = 0, n = strlen(cad); lon = 0;
    if (n==1 && cad[0]=='1')
      printf("1\n");
    else {
      for (int i = 0; i < n; i++){
          int result = (cad[n-i-1]-48)*2 + acarreo; 
          acarreo = (result > 9) ? 1 : 0; 
          solve[lon++] = result%10+48;
      }
      if (acarreo) solve[lon++] = 49; 
      acarreo = 2;
      for (int i = 0; i < lon; i++)
         if ((solve[i]-48)-acarreo >= 0)
              { solve[i] -= acarreo; break; }
         else
              { solve[i] = solve[i]+10-acarreo; acarreo = 1; }
      while (solve[lon-1]=='0') lon--;
      for (int i = lon-1; i >=0; i--) printf("%c", solve[i]);
      printf("\n");
    }
  }
}
