#include <bits/stdc++.h>
using namespace std;

typedef long long ll;

void solve() {

    int t; cin >> t;

    for (int times = 0; times < t; times++) {

        int n; cin >> n;
        char c; cin >> c;

        char a[n];
        for (auto &i : a) { cin >> i; }        

        if (c == 'g') {
            cout << "0" << endl;
            continue;
        }

        int gmax = 0;

        for (int i = 0; i < n; i++) {

            if (a[i] == c) {

                int g = 0;
                for (int j = i+1; j < i+n; j++) {

                    int index;
                    if (j >= n) {
                        index = j - n;
                    } else {
                        index = j;
                    }

                    g++;
                    if (a[index] == 'g') break;
            
                }
                
                if (g > gmax) {
                    gmax = g;
                }

            }
        }

            cout << gmax << endl;
        }


    

}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
} 