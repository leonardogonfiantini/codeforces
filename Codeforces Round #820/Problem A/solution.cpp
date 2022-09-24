#include <bits/stdc++.h>
using namespace std;

void solve() {

    int t; cin >> t;

    for (int i = 0; i < t; i++) {
        int a, b, c;
        cin >> a >> b >> c;

        int bc = abs(b - c);
        int bone = abs(c - 1);
        int bco = c + bc;

        cout << "\n";

        if (bco < a) cout << 2;
        if (bco > a) cout << 1;
        if (bco == a) cout << 3;
    }


}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
}