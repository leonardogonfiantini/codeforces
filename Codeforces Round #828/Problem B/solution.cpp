#include <bits/stdc++.h>
using namespace std;

typedef long long ll;

void solve() {

    int t; cin >> t;

    for (int times = 0; times < t; times++) {

        int n, q; cin >> n >> q;

        ll a[n];
        for (auto &i : a) { cin >> i; }

        ll sum = 0;
        int p = 0;
        int d = 0;
        for (int i = 0; i < n; i++) { 
            
            if ((a[i] % 2) == 0) {
                p++;
            } else {
                d++;
            }

            sum += a[i]; 
        }

        
        for (int i = 0; i < q; i++) {

            int o, x; cin >> o >> x;

            if (o == 0) {
                sum += p*x;

                if (x % 2 != 0) {
                    d += p;
                    p = 0;

                }

            } else {
                sum += d*x;

                if (x % 2 != 0) {
                    p += d;
                    d = 0;

                }
            }

            cout << sum << endl;
        }
    
    }


}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
} 