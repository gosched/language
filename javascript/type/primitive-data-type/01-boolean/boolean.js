console.log(typeof true);

// equality operator
// strict equality operator

// true
    // null == undefined

    // true == 1
    // true == '1'
    // true == "1"

    // false == 0
    // false == '0'
    // false == "0"
    // false == ''
    // false == ""

    // 0 == '0'
    // 0 == "0"

    // " \t\r\n" == 0
    // " \t\r\n" == false

// false
    // console.log(" \t\r\n" == "");
    // console.log(" \t\r\n" == '');
    // console.log(true === 1);
    // console.log(true == 2);

// equal true
    // []
    // {}

// equal false
    // false
    // null
    // undefined
    // NaN
    // ""
    // ''

let flag;
if (flag) {
    console.log('!!!');
}

flag = null;
if (flag) {
    console.log('!!!');
}

flag = false;
if (flag) {
    console.log('!!!');
}

flag = 0;
if (flag) {
    console.log('!!!');
}

flag = NaN;
if (flag) {
    console.log('!!!');
}

flag = "";
if (flag) {
    console.log('!!!');
}

flag = '';
if (flag) {
    console.log('!!!');
}

flag = [];
if (flag) {
    console.log('[] == true');
}

flag = {};
if (flag) {
    console.log('{} == true');
}