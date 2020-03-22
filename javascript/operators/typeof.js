// 確定一個值是什麼類型
// typeof
// instanceof
// Object.prototype.toString

// 確定一個值是否存在
// if (typeof variable === "undefined") {}
// if (property in object) {}

// Type
    // undefined
        // let x;
    // null
        // ?
    // boolean
        // true
        // false
    // number
        // NaN
        // 1
        // 3.14
    // symbol
        //
        // Symbol()
        // Symbol('foo')
    // string
        // ''
        // ""
    // object
        // null
        // []
        // {}
        // new Array()
        // new Map()
    // function
        // Object

// The typeof operator returns a string indicating the type of the unevaluated operand.

// console.log(typeof(undefined));

// console.log(typeof(true));
// console.log(typeof(false));

// console.log(typeof(NaN));
// console.log(isNaN(NaN));

// console.log(typeof(1));
// console.log(typeof('1'));
// console.log(typeof("1"));
// console.log(typeof(new Number(1)));

// console.log(typeof(3.14));
// console.log(typeof('3.14'));
// console.log(typeof("3.14"));
// console.log(typeof(new Number(3.14)));

// const symbol1 = Symbol();
// const symbol2 = Symbol('foo');
// console.log(typeof symbol1);
// console.log(symbol2.toString());
// console.log(Symbol('foo') === Symbol('foo'));

// Array is Object
// let array1 = [];
// let array2 = new Array();
// console.log(typeof(array1));
// console.log(typeof(array2));
// console.log(array1 instanceof Array);
// console.log();

// let map = new Map();
// console.log(typeof(map));
// console.log(typeof({}));
// console.log(typeof(new Object()));

// console.log(typeof(Object));

// 為了兼容
// typeof null == "object"
// console.log(typeof null);

