// Object <string, value>
// Map    <value, value>

// const map = new Map();
// console.log(`map.set('site', 'github.com'): ${map.set('site', 'github.com')}`);
// console.log(`map.get('site'): ${map.get('site')}`);
// console.log(`map.delete('site'): ${map.delete('site')}`);
// console.log(`map.delete('site'): ${map.delete('site')}`);
// console.log(`map.clear(): ${map.clear()}`);
// console.log(`map.size: ${map.size}`);

// map.set('key1', 1);
// map.set('key2', 2);
// map.set('key3', 3);

// for (const [k, v] of map) {
//     console.log(k, v);
// }

let m = new Map();

m.set(1, 123);

console.log(m.get(1)); // 123
console.log(m.get(2)); // undefined
console.log(m.has(1)); // true
console.log(m.size);   // 1

// return true or false
m.delete(1);

m.clear();

console.log();

// Map from Array, Set, Map
const m1 = new Map([
    ['name1', 'javascript'],
    ['name2', 'java']
]);

const set = new Set([
    ['foo', 1],
    ['bar', 2]
]);

const m2 = new Map(set);

const m3 = new Map(m2);
console.log();

const m4 = new Map();

const k1 = ['a'];
const k2 = ['a'];

m4.set(k1, 111).set(k2, 222);

console.log(m4.get(k1)); // 111
console.log(m4.get(k2)); // 222
console.log();

// key1 === key2 -> same key
// NaN  !== NaN  -> same key
let m5 = new Map();

m5.set(-0, 123);
m5.set(true, 1);
m5.set('true', 2);
m5.set(undefined, 3);
m5.set(null, 4);
m5.set(NaN, 123);

console.log(m5.get(+0));        // 123
console.log(m5.get(true));      // 1
console.log(m5.get(undefined)); // 3
console.log(m5.get(null));      // 4
console.log(m5.get(NaN) );      // 123
console.log();

let m6 = new Map([
    [1, 'a'],
    [2, 'b'],
    [3, 'c']
]);

// m.keys();
for (let key of m6.keys()) {
    console.log(key);
}
  
// m.values();
for (let value of m6.values()) {
    console.log(value);
}

// m.entries();
for (let item of m6.entries()) {
    console.log(item[0], item[1]);
}

for (let [key, value] of m6.entries()) {
    console.log(key, value);
}

for (let [key, value] of m6) {
    console.log(key, value);
}

// m[Symbol.iterator] === m.entries

// m.forEach();
m6.forEach(function(value, key) {
    console.log("Key: %s, Value: %s", key, value);
});


const m7 = new Map([
    [1, 'one'],
    [2, 'two'],
    [3, 'three'],
]);

console.log([...m7.keys()]);
console.log([...m7.values()]);
console.log([...m7.entries()]);
console.log([...m7]);

function mapToObjet(map) {
    let obj = Object.create(null);
    for (let [k,v] of map) {
      obj[k] = v;
    }
    return obj;
}

function objToMap(obj) {
    let map = new Map();
    for (let k of Object.keys(obj)) {
        map.set(k, obj[k]);
    }
    return map;
}

// Map to JSON
// Map key type == string
// let json1 = JSON.stringify(mapToObjet(strMap));

// Map key type != string
// let json2 = JSON.stringify([...map]);

// JSON to Map
// objToMap(JSON.parse(jsonStr))