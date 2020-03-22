// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Set

// The Set object lets you store unique values of any type,
// whether primitive values or object references.

const set = new Set();
console.log(set.size);
set.add(1).add(2).add(3);
set.delete(1);
set.has(1);
set.clear();

console.log(set.add(1).add(2).add(3));
console.log(set.delete(1));
console.log(set.has(1));
console.log(set.clear());

// set1.entries();
// set1.keys();
// set1.values();