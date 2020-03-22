// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol/for
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol/keyFor
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/getOwnPropertySymbols

// https://harttle.land/2018/10/14/whats-symbols-for.html

// Symbol
// 非私有屬性
// 類似私有屬性

// 可以獲取 Symbol
// Reflect.ownKeys() 
// Object.getOwnPropertySymbols()

// 無法獲取 Symbol
// JSON.stringify()
// for ... in
// Object.keys()
// Object.getOwnPropertyNames()

// var s0 = new Symbol('test'); // TypeError
var s1 = Symbol('test');
var s2 = Symbol('test');
var s3 = Symbol('hello');
var s4 = Symbol.for('hello');
var s5 = Symbol.for('hello');

console.log(typeof s1);      // symbol
console.log(s1);             // Symbol(test)
console.log(s1.toString());  // Symbol(test)
console.log(s1.valueOf());   // Symbol(test)
console.log(s1.description); // test

console.log(s1 instanceof Symbol); // false
console.log(s1 === s2, s1 == s2);  // false false
console.log(s3 === s4, s3 == s4);  // false false
console.log(s4 === s5, s4 == s5);  // true true

// console.log(Symbol("foo") + "bar");  // TypeError

let string1 = 'string';
let string2 = 'string';
console.log(string1 == string2);
console.log(string1 === string2);
console.log();

let symbol1 = Symbol('description');
let symbol2 = Symbol('description');
let symbol3 = Symbol();
let symbol4 = Symbol();
console.log(symbol1 == symbol2);
console.log(symbol3 == symbol4);
console.log(symbol1 === symbol2);
console.log(symbol3 === symbol4);
console.log(symbol3);
console.log();

let key = Symbol('symbol');
let object = {}
let symbolObject = Object(key);
console.log(typeof symbolObject);

console.log(object.key === object['key']); // true
console.log(object.key === object["key"]); // true
console.log(object.key === object[key]);   // true

object.key = 1;
object['key'] = 2;
object[key] = 3;

console.log(object.key === object['key']); // true
console.log(object.key === object["key"]); // true
console.log(object.key === object[key]);   // false

console.log(object.key);    // 2
console.log(object['key']); // 2
console.log(object[key]);   // 3
console.log();

// The Symbol.for() method searches for existing symbols in a runtime-wide symbol registry
// with the given key and returns it if found.
// Otherwise a new symbol gets created in the global symbol registry with this key.

// The Symbol.keyFor() method retrieves a shared symbol key from the global symbol registry
// for the given symbol.

const globalSymbol = Symbol.for('foo');
const localSymbol = Symbol();
console.log(Symbol.keyFor(globalSymbol));    // foo
console.log(Symbol.keyFor(localSymbol));     // undefined
console.log(Symbol.keyFor(Symbol.iterator)); // undefined
console.log();

// Object.getOwnPropertySymbols()
// The Object.getOwnPropertySymbols() method returns an array of all symbol properties
// found directly upon a given object.

let object1 = {};
const a = Symbol('a');
const b = Symbol.for('b');
object1[a] = 'localSymbol';
object1[b] = 'globalSymbol';
const symbols = Object.getOwnPropertySymbols(object1);
console.log(Reflect.ownKeys(object1));
console.log(symbols);
console.log(symbols.length);