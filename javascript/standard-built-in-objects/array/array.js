// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array

// Array is object
console.log(typeof (new Array()));
console.log(typeof []);

// object's key == string or symbol
let array1 = [1,2,3,4,5];
console.log(Object.keys(array1));
console.log(array1['0']);
console.log(array1[0]);
array1.a = 100;
console.log(Object.keys(array1));

let array2 = new Array(5);
console.log(Object.keys(array2));
console.log(array2['0']);
console.log(array2[0]);
array2[0] = 101;
array2.a = 102;
console.log(array2.a);
console.log(Object.keys(array2));