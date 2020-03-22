// The name of the function.
function f1() {
    console.log('f1');
}

let f2 = function () {
    console.log('f2');
}

let f = function f3() {
    console.log('f3');
};

console.log(f1.name);
console.log(f2.name);
console.log(f.name);

// The length property indicates the number of parameters expected by the function.
function f4(a, b) {
    let c = a + b + a * b + a / b;
    return c;
}
console.log(f4.length);

// Returns a string representing the source code of the function.
console.log(f4.toString());
console.log(Math.sqrt.toString()); // function name() { [native code] }

// The scope of a variable declared with var is its current execution context
var v = 1;
function f5() {
    var v = 2;
    console.log(v); // 2
}

f5();
console.log(v); // 1

// function f6(x) {
//     console.log(temp);
//     if (x > 100) {
//         var temp = x - 100;
//     }
// }

// function f6(x) {
//     var temp;
//     console.log(temp);
//     if (x > 100) {
//       tmp = x - 100;
//     };
// }

// JavaScript 允許函數有不定數目的參數、引數
// 呼叫時省略引數 函數仍可執行
// 呼叫時增加額外引數 函數仍可執行
function f7(a, b) {
    console.log(arguments);
    console.log(a, b);
    return a;
}
f7(1, 2);
f7(1, 2, 3);
f7(3.14);
f7(1, null);
f7(1, undefined);

// Primitive data type
// pass by value
let f8Var = 1;
function f8(f8Var) {
    f8Var = 2;
}
f8(f8Var);
console.log(f8Var); // 1

// Composite data type
// pass by reference
let f9array = [1, 2, 3];
function f9(f8Var) {
    f9array[0] = 100;
}
f9(f9array);
console.log(f9array); // [100,2,3]

// 如果有同名的參數，則取最後出現的那個值
// 即使後面的 a 沒有值或被省略，也是以其為準
// 如果要獲得第一個 a 的值，可以使用 arguments
function f10(a, a) {
    console.log(a, arguments[0], arguments[1]);
}
f10(1, 2); // 2, 1, 2
f10(1);   // undefined 1 undefined

// arguments 對象可以在運行時修改參數的值
// 嚴格模式下修改 arguments 對像不會影響到實際的函數參數
var f11 = function (a, b) {
    // 'use strict';
    arguments[0] = 1;
    return a + b;
}

console.log(f11(0, 0));
// 0 + 0 == 0
// 1 + 0 == 1