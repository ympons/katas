#include <stdio.h>
float n, i;

void Cards(){
   float  valor = 0;
   i = 1;
   while (valor<n){
      i += 1;
      valor += (1/i);
   }
};
int main() {
    scanf("%f",&n);
    while (n!=0.00){
       Cards();
       printf("%0.f card(s)\n",i-1);
       scanf("%f",&n);
    }
    return 0;
}
