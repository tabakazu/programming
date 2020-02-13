
```bash
# compile & run
$ g++-8 -o A-1.00.exe A-1.00.cpp; ./A-1.00.exe

# compile & run script
$ ./run A-1.00.cpp
```

> AtCoder コードテストの使い方
> https://atcoder.jp/contests/apg4b/tasks/APG4b_ak?lang=ja
 
N 人の合計点を求めるプログラムを作る場合、
「入力の個数 N」と「点数を表す N 個の整数」を入力で受け取り、点数の合計を出力

int N; // 入力の個数 N
cin >> N;

int sum = 0; // 合計点を表す変数
int x;       // 入力を受け取る変数
int i = 0;   // カウンタ変数

while (i < N) { // 入力の個数 N よりカウンタ変数が低い場合
  cin >> x; // ループごとに入力を受け取る
  sum += x; // 合計点に追加する
  i++;      // カウントアップ
}

cout << sum << endl; // 出力
