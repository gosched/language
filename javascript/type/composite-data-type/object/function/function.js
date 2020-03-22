// typeof function == function
// typeof class    == function
console.log(typeof function () { });
console.log(typeof class { });

// Function Declaration
function print(s) {
  console.log(s);
};

// Function Expression
let add = function (x, y) {
  return x + y;
};

// Function Expression
let sub = new Function(
  'x',
  'y',
  'return x - y'
);

// Hoisting
function f() {
  console.log(1);
}
f() // 2

function f() {
  console.log(2);
}
f() // 2

// Return undefined
let result = f();
console.log(result); // undefined

// Recursion