// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/isNaN

// NaN
// 0 除以 0
// 將字符串解析成數字出錯
// 一些數學函數的運算結果
console.log(typeof NaN);    // number
console.log(NaN === NaN);   // false
console.log(Boolean(NaN));  // false
console.log(NaN + 32);      // NaN
console.log(0 / 0);         // NaN
console.log(5 - 'x');       // NaN
console.log(Math.sqrt(-1)); // NaN
console.log(Math.log(-1));  // NaN

// isNaN()
console.log(isNaN(NaN));         // true
console.log(isNaN('Hello'));     // true
console.log(isNaN({}));          // true
console.log(isNaN([123, 321]));  // true
console.log(isNaN([123]));       // false
console.log(isNaN([]));          // false

function typeOfNaN(x) {
    if (Number.isNaN(x)) {
        return 'Number NaN';
    }
    if (isNaN(x)) {
        return 'NaN';
    }
}

console.log(typeOfNaN('100'));
console.log(typeOfNaN('10F'));
console.log(typeOfNaN(NaN));
