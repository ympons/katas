#include <stdio.h>
int m , n , mod, god;

int main() {
	int t, aux;
	scanf("%d", &t);
	while (t--) {
		scanf("%d%d", &n, &m);
		mod = 0; god = 0;
		for (int i = 0; i < n; i++) {
			scanf("%d", &aux);
			if (aux  > god) god = aux;
		}
		for (int i = 0; i < m; i++) {
			scanf("%d", &aux);
			if (aux  > mod) mod = aux;
		}
		if (god >= mod)
			printf("Godzilla\n");
		else
			printf("MechaGodzilla\n");		
	}
	return 0;
}
