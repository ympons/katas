#include <stdio.h>

int k, m, cursos[101];

int main() {
    int curso, cat, kp; bool puedo;
    scanf("%d", &k);
    while (k!=0) {
        scanf("%d", &m);
        for (int i = 0; i < k; i++) scanf("%d", &cursos[i]);
        puedo = true;
        for (int i = 0; i < m; i++) {
            scanf("%d%d", &kp, &cat); int cont = 0;
			while (kp--) {
				scanf("%d", &curso);
				for (int j = 0; j < k; j++)
					if (curso==cursos[j]) { cont++; break; }
			}
			if (cont < cat) puedo = false;
		}
		if (puedo)
			printf("yes\n");
		else
			printf("no\n");
    	scanf("%d", &k);
	}
	return 0;
}
