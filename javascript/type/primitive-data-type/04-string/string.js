console.log(typeof '');

// 字符串可以被視為字符數組
var s = 'vscode';
console.log(s[0]);
console.log('javascript'[0]);

// length 屬性返回字符串的長度
console.log(s.length);

// 如果方括號中的數字超過字符串的長度 返回 undefined
console.log(s[100]);

// 無法改變字符串之中的字符
// 改變會失敗
delete s[0];
s[1] = 'abc';
console.log(s);

// JavaScript 引擎內部 所有字符都用 Unicode 表示
// 會自動識別字符是字面形式，還是 Unicode 形式
// 輸出給用戶的時候，所有字符都會轉成字面形式
var s = '\u00A9';
console.log(s);

var f\u006F\u006F = 'abc';
console.log(foo); // abc

// 每個字符在 JavaScript 都是以 UTF-16 格式儲存
// JavaScript 對 UTF-16 的支持是不完整的
// JavaScript 支持 16 bits（2 bytes）的字符

// UTF-16 有兩種
// 對於 code point 在 U+0000 到 U+FFFF 之間的字符，長度為 16 bits（2 bytes）， length == 1

// 對於 code point 在 U+10000 到 U+10FFFF 之間的字符，長度為 32 bits（4 bytes）， length == 2
// 前兩個字節在 0xD800 到 0xDBFF 之間，後兩個字節在 0xDC00 到 0xDFFF 之間

// JavaScript 返回的字符串長度可能是不正確的
// 舉例來說，code point U+1D306 對應的字符為 𝌆，它寫成 UTF-16 是 0xD834 0xDF06
console.log('𝌆'.length); // 2

// Base64 是一種編碼方法
// 可以將任意值轉成 0～9 A～Z a-z + 和 / 這64個字符組成的可打印字符串
// 使用它的主要目的，是為了不出現特殊字符，簡化程序的處理，不是為了加密
console.log(Buffer.from("Hello World").toString('base64'));
console.log(Buffer.from("SGVsbG8gV29ybGQ=", 'base64').toString('ascii'));

// 轉義特殊字符
/**
    \0 ：null
    \b ：後退鍵
    \f ：換頁符
    \n ：換行符
    \t ：製表符
    \v ：垂直製表符
    \' ：單引號
    \" ：雙引號
    \\ ：反斜杠
*/