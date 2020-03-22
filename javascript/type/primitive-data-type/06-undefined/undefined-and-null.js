console.log(null === undefined); // false
console.log(null == undefined); // true
console.log(null == false); // false
console.log(!null == true); // true
console.log(Number(null)); // 0
console.log(Number(undefined)); // NaN (not a number)
console.log();

// null 表示空值
// undefined 表示未定義

var x; // x == undefined

function f1(data) { return data } // data == undefined
console.log(f1()); // undefined

function f2() { }
console.log(f2()); // undefined