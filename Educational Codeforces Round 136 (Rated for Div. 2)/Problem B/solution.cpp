#include <bits/stdc++.h>
using namespace std;

void solve() {

    int t; cin >> t;


    for (int i = 0; i < t; i++) {

        int n; cin >> n;

        int arr[n];
        for (auto &i : arr) { cin >> i; }

        int corr[n];

        corr[0] = arr[0];

        int f = 0;

        for (int j = 1; j < n; j++) {

            int val1 = arr[j] + corr[j-1];
            int val2 = corr[j-1] - arr[j];

            if ((val1 - corr[j-1]) == arr[j] && (corr[j-1] - val2) == arr[j]) {
                cout << "-1" << endl;
                f = 1;
                break;
            }

            corr[j] = val1;

        }

        if (f == 0) {
            for (auto i : corr) { cout << i << " "; }
            cout << endl;
        }



    } 


}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
}