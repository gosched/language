// javascript 的 length 類似 go 的 capacity
// 空位不在尾端會影響 length 屬性的值
// 空位的預設值為 undefined
// 使用 delete 刪除一個數組成員 會形成空位 不會影響 length

let a1 = [1, , 3];
console.log(a1[1]); // undefined
console.log(a1.length); // 3

let a2 = [3, 2, 1,];
console.log(a2.length); // 3
delete a2[0];
delete a2[1];
console.log(a2); // [ <2 empty items>, 1 ]
console.log(a2[0]); // undefined
console.log(a2.length); // 3
console.log();

// for (let i = 0; i < length; i++)、for...of
// 如果某個位置是 undefined 或 空位 遍歷的時候不會被跳過
let array = [1, 2, undefined, 4, , 6];
for (let i = 0; i < array.length; i++) {
    // if (array[i])
    console.log(array[i]);
}
for (let value of array) {
    console.log(value);
}
console.log();

// Object.keys、for...in、forEach
// 如果某個位置是 undefined 遍歷的時候不會被跳過
// 如果某個位置是 空位 遍歷的時候會被跳過
let array2 = [, , , , undefined];
console.log(Object.keys(array2));
for (let i in array2) {
    console.log(i);
}
array2.forEach(function (e) {
    console.log(e);
});