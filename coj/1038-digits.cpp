#include <iostream>
using namespace std;

string cad1, cad0; int pos;

string toStr(int i) {
   string num; int temp;
   while(i/10!=0) {
       temp  = i%10; i /= 10;
       temp += 48;
       num = (char)temp + num;
   }
   i += 48;
   num = (char)i + num ;
   return num;
}

int main() {
    cin>>cad1;
    while (cad1!="END") {
        pos  = 1;
        cad0 = toStr(cad1.length());
        while (cad1!=cad0) {
            pos++;
            cad1 = cad0;
            cad0 = toStr(cad1.length());
        }
        cout<<pos<<endl;
        cin>>cad1;
    }
    return 0;
}
