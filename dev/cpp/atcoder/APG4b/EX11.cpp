#include <bits/stdc++.h>
using namespace std;

int main() {
  int N, A;
  cin >> N >> A;

  int total = A;
  for (int i = 0; i < N; i++) {
    string op;
    int B;
    cin >> op >> B;

    if (op == "+") {
      total += B;
    } else if (op == "-") {
      total -= B;
    } else if (op == "*") {
      total *= B;
    } else if (op == "/" && B != 0) {
      total /= B;
    } else {
      cout << "error" << endl;
      break;
    }

    cout << i + 1 << ":" << total << endl;
  }
}
