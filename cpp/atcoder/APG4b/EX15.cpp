#include <bits/stdc++.h>
using namespace std;

int sum(int n) {
  int total = 0;
  for (int i = 0; i < n; i++) {
    int X;
    cin >> X;
    total += X;
  }
  return total;
}

int main() {
  int N;
  cin >> N;

  int A = sum(N);
  int B = sum(N);
  int C = sum(N);

  cout << A * B * C << endl;
}
