#include <bits/stdc++.h>
using namespace std;

typedef long long ll;

void solve() {

    int t; cin >> t;

    for (int times = 0; times < t; times++) {

        ll n; cin >>n;

        char dir[n];
        for (auto &i : dir) {cin >> i;}

        int f = 0;

        for (ll i = 1; i <= n; i++) {
            ll l = 0;
            ll r = n-1;
            ll k = i;
            ll val = 0;


            if (f == 0) {
                
                while (l <= r) {


                    if (l == r) {
                        val += n/2;
                    } else {

                        if (dir[l] == 'L') {
                            if (k > 0) {k--; val += n - l - 1;}
                            else {val += l;}
                        } else {
                            val += n - l - 1;
                        }

                        if (dir[r] == 'R') {
                            if (k > 0) {k--; val += r;}
                            else {val += n - r - 1;}
                        } else {
                            val += r;
                        }
                    }
                    

                    l++;
                    r--;
                }
            }

            if (k > 1 && val > 0) {
                f = val;
            }
            
            if (f > 0) {
                cout << f << " ";
            } else {
                cout << val << " ";
            }
        }

        cout << endl;
    }
}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
}