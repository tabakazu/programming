#include <bits/stdc++.h>
using namespace std;

int main() {
  int p;
  cin >> p;
  string text;

  if (p == 2) {
    cin >> text;
  }

  int price;
  cin >> price;
  int N;
  cin >> N;

  if (p == 2) {
    cout << text << "!" << endl;
  }
  cout << price * N << endl;
}
