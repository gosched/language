let array1 = [1,2,3,4,5];

let slice1 = array1.slice(1);    // [ 2, 3, 4, 5 ]
let slice2 = array1.slice(0, 3); // [ 1, 2, 3 ]
let slice3 = array1.slice(1, 7); // [ 2, 3, 4, 5 ]
let slice4 = array1.slice(4);    // [ 5 ]
let slice5 = array1.slice(10);   // []
let slice6 = array1.slice(-1);   // [ 5 ]
let slice7 = array1.slice(-3);   // [ 3, 4, 5 ]

console.log(array1);
console.log();

console.log(slice1);
console.log(slice2);
console.log(slice3);
console.log(slice4);
console.log(slice5);
console.log(slice6);
console.log(slice7);
console.log();

console.log(array1);
console.log();