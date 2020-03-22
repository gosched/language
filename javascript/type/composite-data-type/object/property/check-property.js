// in
// for in
// hasOwnProperty
// Object.keys

let o1 = {
  abc: 123
};

console.log('toString' in o1);
if ('toString' in o1) {
  console.log(o1.hasOwnProperty('toString')) // false
}

// for...in
// 不僅遍歷對象自身的屬性 還遍歷繼承的屬性
// 遍歷 enumerable 的屬性 跳過不可遍歷的屬性
for (let key in o1) {
  console.log(key, o1[key]);
}

let user = { name: 'vscode' };
for (let key in user) {
  if (user.hasOwnProperty(key)) {
    console.log(key);
  }
}

// Returns the names of the enumerable string properties and methods of an object.
console.log(Object.keys(o1)); // [ 'abc' ]