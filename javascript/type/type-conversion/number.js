// Number 比 parseInt 嚴格
// 當 Number 的參數是 object 時 將返回 NaN 除非參數是包含單個數值的陣列
// parseInt 和 Number 都會過濾前綴後綴的空格

let n1 = Number(324)       // 324

let n2 = Number(undefined) // NaN
let n3 = Number(null)      // 0

let n4 = Number(false)     // 0
let n5 = Number(true)      // 1

let n6 = Number([])        // 0
let n7 = Number([1])       // 1
let n8 = Number([2])       // 2
let n9 = Number(new Map()) // NaN

let n10 = Number('')                // 0
let n11 = Number('324')             // 324
let n12 = Number('324abc')          // NaN 
let n13 = Number('\t\v\r12.34\n')   // 12.34
let n14 = parseInt('\t\v\r12.34\n') // 12

console.log(n1);
console.log(n2);
console.log(n3);
console.log(n4);
console.log(n5);
console.log(n6);
console.log(n7);
console.log(n8);
console.log(n9);
console.log(n10);
console.log(n11);
console.log(n12);
console.log(n13);
console.log(n14);
