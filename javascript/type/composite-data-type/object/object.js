// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Object_initializer

// Objects can be initialized using new Object(), Object.create(), or using the literal notation (initializer notation).
// An object initializer is a comma-delimited list of zero or more pairs of property names
// and associated values of an object, enclosed in curly braces ({}).

// 對像是一組 key-value 的集合
// 所有鍵名都是 string or symbol

// console.log(typeof (new Object()));
// console.log(typeof console);
// console.log(typeof process);

// let o = {
// foo:'hello',
// bar:'world'
// };

// o.test = '!';
// console.log(o.foo);
// console.log(o['foo']);
// console.log(Object.keys(o));
// console.log(o);
// console.log();

// 對象的每一個 key 又稱為 property
// property 的 value 可以是任何數據類型
// 如果一個 property 的值為 function 通常把這個 property 稱為 method
// let math = {
// pow: function (x) {
// return x * x;
// }
// };

// console.log(math.pow(2)); // 4
// console.log();

// 如果屬性的值還是一個對象 就形成了鍊式引用
// let foo = {};
// let bar = {
// hello:'world'
// };
// foo.bar = bar;
// console.log(foo.bar.hello);
// console.log();

// console.log(Object.name); // Object
// console.log(Object.length); // 1
// console.log(Object.prototype); // {}
// console.log(Object.prototype.constructor); // [Function: Object]
// console.log(Object.caller);
// console.log(Object.arguments);
// console.log();

// Object.setPrototypeOf();
// Object.getPrototypeOf();

// Object.defineProperty();
// Object.defineProperties();
// Object.getOwnPropertyNames();
// Object.getOwnPropertySymbols();
// Object.getOwnPropertyDescriptor();
// Object.keys();

// Object.toString();

// Object.is();

// Object.isExtensible();
// Object.isFrozen();
// Object.isSealed();

// Object.preventExtensions();
// Object.freeze();
// Object.seal();

// Object.create();
// Object.apply();
// Object.assign();
// Object.bind();
// Object.call();

// The Object constructor creates an object wrapper for the given value.
// If the value is null or undefined, it will create and return an empty object,
// otherwise, it will return an object of a Type that corresponds to the given value.
// If the value is an object already, it will return the value.

// When called in a non-constructor context, Object behaves identically to new Object().

// var o1 = new Object();
// var o2 = new Object(undefined);
// var o3 = new Object(null);

// var o4 = new Object(true); // equivalent to o = new Boolean(true);
// var o5 = new Object(Boolean()); // equivalent to o = new Boolean(false);

// var o6 = new Object('123');
// var o7 = new Object([1,2,3]);

// console.log(typeof (o5), o5, Object.keys(o5)); // object [Boolean: false] []
// console.log(typeof (o6), o6, Object.keys(o6)); // object [String: '123'] [ '0', '1', '2' ]
// console.log(typeof (o7), o7, Object.keys(o7)); // object [ 1, 2, 3 ] [ '0', '1', '2' ]