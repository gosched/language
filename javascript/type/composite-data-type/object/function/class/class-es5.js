function Point(x, y, z) {
    this.x = x;  // public
    this.y = y;  // public
    var _z = z;   // private

    this.getZ = function () {
        return _z;
    }
}

let p1 = new Point(3, 3, 4);

console.log(typeof Point); // function
console.log(typeof p1);    // object
console.log(p1);           // Point { x: 3, y: 3, getZ: [Function] }
console.log(p1.getZ());

// Point.prototype.toString = function () {
//     return '(' + this.x + ', ' + this.y + ')';
// };

console.log(p1.toString()); // (3, 3)