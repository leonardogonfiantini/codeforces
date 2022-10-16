#include <bits/stdc++.h>
using namespace std;

void solve() {

    int t; cin >> t;

    for (int times = 0; times < t; times++) {

        int n; cin >> n;

        int a = 0;
        for (int i = 0; i < n; i++) {
            int x; cin >> x;

            if (x == 1) { a++; }
        }

        int b = 0;
        for (int i = 0; i < n; i++) {
            int x; cin >> x;

            if (x == 1) { b++; }
        }

        if (a == b && a == n) {
            cout << "0" << endl;
        } else {

            if (b > a) {
                cout << b-a+1 << endl;
            } else {
                cout << a-b << endl;
            }

        }

    }

}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
}