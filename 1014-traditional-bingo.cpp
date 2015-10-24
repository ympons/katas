#include <stdio.h>

int jugadas, numero, col;
bool isbingo;

int main() {
    int t; scanf("%d",&t);

    while(t--) {
        int m[6][6]; isbingo = false;  jugadas = 0; m[3][3] = 0;
        scanf("%d%d%d%d%d",&m[1][1],&m[1][2],&m[1][3],&m[1][4],&m[1][5]);
        scanf("%d%d%d%d%d",&m[2][1],&m[2][2],&m[2][3],&m[2][4],&m[2][5]);
        scanf("%d%d%d%d",  &m[3][1],&m[3][2],&m[3][4],&m[3][5]);
        scanf("%d%d%d%d%d",&m[4][1],&m[4][2],&m[4][3],&m[4][4],&m[4][5]);
        scanf("%d%d%d%d%d",&m[5][1],&m[5][2],&m[5][3],&m[5][4],&m[5][5]);

        for(int i = 0; i < 75; i++) {
            scanf("%d",&numero); col = (numero-1)/15 + 1;
            if(!isbingo) {
                jugadas++;
                for (int j = 1; j < 6; j++) {
                    if (m[j][col] == numero) {
                        m[j][col] = 0;
                        isbingo = ((!m[j][1]   && !m[j][2]   && !m[j][3]   && !m[j][4]   && !m[j][5])   ||
                                   (!m[1][col] && !m[2][col] && !m[3][col] && !m[4][col] && !m[5][col]) ||
                                   (!m[1][1]   && !m[2][2]   && !m[3][3]   && !m[4][4]   && !m[5][5])   ||
                                   (!m[1][5]   && !m[2][4]   && !m[3][3]   && !m[4][2]   && !m[5][1]));
                        if (isbingo == true) break;
                     }
                }
            }
        }
        printf("BINGO after %d numbers announced\n",jugadas);
    }
    return 0;
}
