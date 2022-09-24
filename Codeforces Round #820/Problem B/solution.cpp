#include <bits/stdc++.h>
using namespace std;

void solve() {

    int t; cin >> t;

    for (int i = 0; i < t; i++) {

        int n; cin >> n;
        
        char arr[n];
        for (int j = 0; j < n; j++) {
            char x; cin >> x;
            arr[j] = x;
        }


        string s = "";
    
       for (int j = n-1; j >= 0; j--) {
            if (arr[j] != '0') {
                s = (char)(48 + arr[j]) + s;
            } else {
                s = (char)(48 + (arr[j-2] - '0')*10 + arr[j-1]) + s;
                j -= 2;
            }
        }

        cout << s << "\n";


    }


}

int main() {
    cin.tie(0)->sync_with_stdio(0);
    int T = 1;
    while(T--) solve();
}