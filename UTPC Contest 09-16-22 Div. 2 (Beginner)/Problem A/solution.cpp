#include <bits/stdc++.h>
using namespace std;

void solve() {

    int k = 0; cin >> k;

    vector<int> v;

    int x;
    while (cin >> x) {
        v.push_back(x);
    }

    sort(v.begin(), v.end());

    cout << v[(int) v.size() - k] << "\n";

}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
}