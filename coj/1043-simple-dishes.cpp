#include <stdio.h>
const int l[15] = {16384,8192,4096,2048,1024,512,256,128,64,32,16,8,4,2,1};
int h, n, v[1000], i, j, c[100000][4] = {0};
int main() {
  c[0][3] = 1;
  for (i = 0; i < 15; i++)
    for (j = 32767 - l[i]; j >= 0; j--)
      if (c[j][3] != 0)
        if (c[j + l[i]][3] < c[j][3] + l[i]) {
            c[j + l[i]][0] = i;
            c[j + l[i]][1] = j;
            c[j + l[i]][2] = c[j][2] + 1;
            c[j + l[i]][3] = c[j][3] + l[i];
        }
	scanf("%d",&n);
	while(n--) {
		scanf("%d",&h);
		j = h;  v[0] = c[h][2];
		for (i = 1; i <= c[h][2]; i++) {
			v[i] = c[j][0]; j = c[j][1];
		}
		printf("%d", 14-v[1]);
		for (i = 2; i <= v[0]; i++) printf(" %d",14-v[i]);
		printf("\n");
	}
 	return 0;
}
