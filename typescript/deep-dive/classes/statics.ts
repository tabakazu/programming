// 静的メンバ

class Something {
  static instances = 0;
  constructor() {
      Something.instances++;
  }
}

var s1: Something = new Something();
console.log(Something.instances);
var s2: Something = new Something();
console.log(Something.instances);
var s3: Something = new Something();
console.log(Something.instances);

// 出力
// 1
// 2
// 3
