#include <bits/stdc++.h>
using namespace std;

void solve() {


    int t; cin >> t;


    for (int times = 0; times < t; times++) {

        int n; cin >> n;
        int a[n];
        for (auto &i : a) { cin >> i; }
        
        string s;
        cin >> s;   

        char w[n];

        for (int i = 0; i < n; i++) {

            char l = s[i];
            int pos = a[i];

            for (int j = 0; j < n; j++) {

                if (a[j] == pos) {
                    w[j] = l;
                }

            }
        }

        int f = 0;
        for (int i = 0; i < n; i++) {
            if (w[i] != s[i]) {
                cout << "NO" << endl;
                f= 1;
                break;
            }
        }

        if (f == 0) {
        cout << "YES" << endl;
        }


        

    }


}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
} 