// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for...in
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for...of

let object = [10, 20, 30];
object.test = 101;
for (let property in object) {
    process.stdout.write(property + ' ' + object[property]+ ' ');
}
console.log(object);
console.log();

let iterable = [10, 20, 30];
for (let value of iterable) {
  value += 1;
  console.log(value);
}
console.log(iterable);
console.log();

let array = [10, 20, 30];
array.forEach(function (element) {
    console.log(++element);
});
console.log(array);
console.log();

for (let i = 0; i < array.length; i++) {
    array[i]++;
    console.log(array[i]);
}
console.log(array);
console.log();

// var l = array.length;
// while (l--) {
//     console.log(l, array[l]);
// }
// console.log();

// var l = array.length;
// while (--l) {
//     console.log(l, array[l]);
// }
// console.log();