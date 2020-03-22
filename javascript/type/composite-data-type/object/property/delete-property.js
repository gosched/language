// delete 一個不存在的屬性 返回 true
let o1 = {};
let deleted1 = delete o1.abc;
console.log(deleted1); // true

// delete 繼承的屬性 無法刪除 返回 true
let o2 = {};
let deleted2 = delete o2.toString;
console.log(deleted2); // true
console.log(o2.toString); // 該屬性沒有被刪除 依然存在

// delete 一個存在且不得刪除的屬性 返回 false
var o3 = {};
Object.defineProperty(o3, 'abc', {
    value: 123,
    configurable: false
});
let deleted3 = delete o3.abc;
console.log(deleted3); // false