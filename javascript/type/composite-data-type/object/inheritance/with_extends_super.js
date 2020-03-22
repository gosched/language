class Point {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }
}

class ColorPoint extends Point {
    constructor(x, y, color) {
        super(x, y);
        this.color = color;
    }
}

class TestPoint extends Point {
    constructor(x, y, color) {

    }
}

let point1 = new ColorPoint(0, 100, 'red');
console.log(point1);

try {
    // 子類沒有呼叫 super 會引發 ReferenceError
    let point2 = new TestPoint(0, 100, 'red');
    console.log(point2);
} catch (error) {
    console.log(error);
}