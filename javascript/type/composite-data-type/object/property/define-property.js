// 'use strict';

const object1 = {};

Object.defineProperty(object1, 'property1', {
    value: 100,
    writable: false
});

object1.property1 = 101;
// throws an error in strict mode

console.log(object1.property1);
// expected output: 100
