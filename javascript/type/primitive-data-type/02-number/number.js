// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number
// https://en.wikipedia.org/wiki/Single-precision_floating-point_format
// https://en.wikipedia.org/wiki/Double-precision_floating-point_format
// http://www.ruanyifeng.com/blog/2010/06/ieee_floating-point_representation.html

console.log(typeof 1);    // number
console.log(typeof 3.14); // number

// JavaScript 沒有整數 所有數字都是小數 (64位浮點數
// JavaScript 運算整數時 會把浮點數轉成整數 再進行運算
// 浮點數不是精確的值 所以涉及小數的比較和運算要特別小心

// 根據國際標準 IEEE 754
// JavaScript 浮點數的 64 個二進制位

// 第 1 位：符號位，0 表示正數，1 表示負數
// 第 2 位 ~ 第 12 位（共11位）：指數部分
// 第 13 位 ~ 第 64 位（共52位）：小數部分（即有效數字）

// 符號 決定了一個數的正負
// 小數 fraction 決定了數值的精度
// 指數 exponent 決定了數值的大小

// +-1 * 1.fraction * 2^(exponent - 1023)

// Number.MAX_SAFE_INTEGER  // 可以精確運算的最大整數
// Number.MIN_SAFE_INTEGER  // 可以精確運算的最小整數
// Number.MAX_VALUE         // 可以表示的最大值
// Number.MIN_VALUE         // 可以表示的最小值
// Number.POSITIVE_INFINITY // Infinity
// Number.NEGATIVE_INFINITY // -Infinity
// Number.NaN               // Not a Number
// Number.EPSILON
// Number.prototype

// [-(2^53 - 1) ~ (2 ^ 53 - 1) ]
// [-9007199254740991 ~ 9007199254740991]

console.log(1.0 === 1);                   // true
console.log(2.0 === 2);                   // true
console.log(0.1 + 0.2 === 0.3);           // false
console.log((0.3 - 0.2) === (0.2 - 0.1)); // false
console.log(0.3 / 0.1);

console.log(Number.MAX_SAFE_INTEGER); //  9007199254740991
console.log(Number.MIN_SAFE_INTEGER); // -9007199254740991
console.log(Math.pow(2, 53) - 1);     //  9007199254740991
console.log(-(Math.pow(2, 53) - 1));  // -9007199254740991

console.log(Math.pow(2, 53) === Math.pow(2, 53) + 1); // true
console.log(Math.pow(2, 53));             // 9007199254740992
console.log(Number.MAX_SAFE_INTEGER + 1); // 9007199254740992
console.log(Number.MAX_SAFE_INTEGER + 2); // 9007199254740992
console.log(Number.MAX_SAFE_INTEGER + 3); // 9007199254740994

// JavaScript 可以精確處理15位的十進制數 (2^53 是一個16位的十進制數值)

// 根據標準 64 位浮點數的指數部分的長度是 11 個二進制位
// 意味著指數部分的最大值是2047（2^11 - 1）
// 指數部分能夠表示的數值範圍為 2^1024 到 2^-1023
console.log(Math.pow(2, 1024));  // Infinity // overflow
console.log(Math.pow(2, -1075)); // 0        // underflow (指數部分-1023，小數部分52位)

// 數值有多種表示方法 可以用字面形式直接表示 比如 十六進制 0xFF == 十進制 255
// 也可以採用科學計數法表示
// 123e3 // 123000
// 123e-3 // 0.123
// -3.1E+12
// .1e-23

// JavaScript 會自動將數值轉為科學計數法表示 (兩種情況)
// 小數點前的數字多於21位
// 小數點後的零多於5個
// console.log(1234567890123456789012);
// console.log(123456789012345678901);
// console.log(0.0000003);
// console.log(0.000003);

// JavaScript 對整數提供四種進制的表示方法
// 十進制、十六進制、八進制、二進制
// 123
// 0x or 0X
// 0o or 0O or 0...(08x 09x not included)
// 0b or 0B
// console.log(0777); // 511
// console.log(0888); // 888

// Infinity, -Infinity
// 正的數值太大 or 負的數值太小 or 非0數值除以0

// Infinity 大於一切數值（除了NaN）
// -Infinity 小於一切數值（除了NaN）
// Infinity 與 NaN 比較 總是返回 false

console.log(Number(null))     // 0
console.log((+0).toString()); // 0
console.log((-0).toString()); // 0
console.log(0 === +0);        // true
console.log(0 === -0);        // true
console.log(+0 === -0);       // true
console.log((1 / +0));        // Infinity
console.log((1 / -0));        // -Infinity

// 某個值是否為正常的數值
// isFinite(Infinity)  // false
// isFinite(-Infinity) // false
// isFinite(NaN)       // false
// isFinite(undefined) // false
// isFinite(null)      // true
// isFinite(-1)        // true

// 字符串轉為整數
console.log(parseInt('100'));
console.log(parseInt('100', 2));

// 字符串轉為浮點數
// parseFloat();

let i = 1;
// i.toExponential()
// i.toFixed()
// i.toLocaleString()
// i.toPrecision()
// i.toString()
// i.valueOf()