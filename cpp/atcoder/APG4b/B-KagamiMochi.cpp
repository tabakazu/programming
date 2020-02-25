#include <bits/stdc++.h>
using namespace std;

int main() {
  int N;
  cin >> N;

  vector<int> vec(N);
  int sum = 0;

  for (int i = 0; i < N; i++) {
    cin >> vec.at(i);
  }

  sort(vec.begin(), vec.end());
  reverse(vec.begin(), vec.end());

  for (int i = 0; i < N; i++) {
    if (i == 0 || vec.at(i - 1) > vec.at(i)) {
      sum = sum + 1;
    }
  }
  cout << sum << endl;
}
