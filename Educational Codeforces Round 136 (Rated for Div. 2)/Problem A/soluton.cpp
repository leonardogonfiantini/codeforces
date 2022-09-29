#include <bits/stdc++.h>
using namespace std;

void solve() {

    int t; cin >> t;

    for (int i = 0; i < t; i++) {

        int n, m;
        cin >> n >> m;

        int f = 0;
        

        for (int r = 1; r <= n; r++) {

            for (int c = 1; c <= m; c++) {

                if  (r >= 2 && c >= 2) {
                    if ((r-2 <= 0 || r-1 <= 0 || r+2 >= n || r+1 >= n) &&
                        (c-2 <= 0 || c-1 <= 0 || c+1 >= m || c+2 >= m)) {

                            cout << r << " " << c << "\n";
                            f = 1;

                            break;


                        }
                }
            }

            if (f == 1) { break;}


        }


        if (f == 0) { cout << n << " " << m << "\n"; }

    }

}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
}