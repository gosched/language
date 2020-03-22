// let a = 1;
// let b = 2;
// let c = 3;

// let a = 1, b = 2, c = 3;
// console.log(a, b, c);

let [a, b, c] = [1, 2, 3];
console.log(a, b, c);

let [item1, item2, item3] = [1, 2];
console.log(item3); // undefined

let [, , data] = ["foo", "bar", "baz"];
console.log(data); // baz

// ...tail == rest parameter
let [head, ...tail] = [1, 2, 3, 4];
console.log(head); // 1
console.log(tail); // [ 2, 3, 4 ]

// ...z == rest parameter
let [x, y, ...z] = ['a'];
console.log(x); // a
console.log(y); // undefined
console.log(z); // []

let [i, j] = [1, 2, 3];
console.log(i); // 1
console.log(j); // 2

let [x1, x2, x3] = new Set(['a', 'b', 'c']);
console.log(x1); // a
console.log(x2); // b
console.log(x3); // c

function* fibs() {
  let a = 0;
  let b = 1;
  while (true) {
    yield a;
    [a, b] = [b, a + b];
  }
}

let [first, second, third, fourth, fifth, sixth] = fibs();
console.log(first);   // 0
console.log(second);  // 1
console.log(third);   // 1
console.log(fourth);  // 2

// 解構賦值可以指定默認值
let [test1 = 1] = [undefined];
console.log(test1); // 1

// 默認值可以引用其他變量
let [a1 = 1, a2 = a1 + 1] = [];
console.log(a2); // 2

// 解構 object
// 屬性名對應變量名
let { bar, foo } = { foo: '123', bar: '12345' };
console.log(foo); // 123
console.log(bar); // 12345

let { foo: baz } = { foo: '2019', bar: '2020' };
console.log(baz); // 2019