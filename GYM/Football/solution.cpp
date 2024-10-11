#include <bits/stdc++.h>

using namespace std;


int main() {

    string players;
    cin >> players;


    int counter = 1;
    char prev_p = players[0];
    for (int i = 1; i < players.length(); i++) {

        if (prev_p == players[i]) {
            counter++;
        } else {
            counter = 1;
        }

        if (counter == 7) {
            break;
        }

        prev_p = players[i];
    }

    if (counter == 7) {
        cout << "YES";
    } else {
        cout << "NO";
    }

}