#include <bits/stdc++.h>
using namespace std;

void solve() {

    int n, c;
    cin >> n >> c;

    vector<int> a(n);
    for (auto &i : a) cin >> i;

    int ans = 0;
    int temp = 1;
    for (int i = 0; i < n; i++) {
        if (i+1 == n || a[i] == a[i + 1]) {
            if (temp > ans) { ans = temp; }
            temp = 0;
        }
        temp++;
    }

    cout << ans << "\n";

}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
}