#include <bits/stdc++.h>
using namespace std;

int main() {
  int N;
  cin >> N;
  int X = N;
  int sum = 0;
  while (N > 0) {
    sum += N % 10;
    N = N / 10;
  }
  if (X % sum == 0) {
    cout << "Yes" << endl;
  } else {
    cout << "No" << endl;
  }
}
