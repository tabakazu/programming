// 継承

class Point {
  x: number;
  y: number;
  constructor(x: number, y: number) {
      this.x = x;
      this.y = y;
  }
  add(point: Point) {
      return new Point(this.x + point.x, this.y + point.y);
  }
}

class Point3D extends Point {
  z: number;
  constructor(x: number, y: number, z: number) {
      super(x, y);
      this.z = z;
  }
  add(point: Point3D) {
      var point2D = super.add(point);
      return new Point3D(point2D.x, point2D.y, this.z + point.z);
  }
}

var p1 = new Point3D(0, 10, 20);
var p2 = new Point3D(10, 20, 30);
var p3 = p1.add(p2);

console.log(p1);
console.log(p2);
console.log(p3);

// 出力
// Point3D { x: 0, y: 10, z: 20 }
// Point3D { x: 10, y: 20, z: 30 }
// Point3D { x: 10, y: 30, z: 50 }
