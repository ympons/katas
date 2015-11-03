#include <stdio.h>
int k,g[102],i,sol;

void QSort(int izq,int der) {
   int i=izq,j=der,v=g[(i+j)/2],temp;
   do{
     while(g[i]<v)i++; while(g[j]>v)j--;
     if(i<=j){temp=g[i];g[i]=g[j];g[j]=temp;i++;j--;}
   }while(i<=j);
   if(izq<j) QSort(izq,j);
   if(der>i) QSort(i,der);
}

int main() {
  scanf("%d",&k);
  for(i=0;i<k;i++) scanf("%d",&g[i]);
  QSort(0,k-1);
  sol=0;
  for(i=0;i<k/2+1;i++) sol+=g[i]/2+1;
  printf("%d",sol);
  return 0;  
}
