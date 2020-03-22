// https://wangdoc.com/javascript/features/error.html

function f1() {
    throw new Error();
}

function f2() {
    try {
        console.log(0);
        f1();
    } catch (e) {
        console.log(1);
        console.log(e);
        return false;
        console.log(3);
    } finally {
        console.log(2);
    }

    console.log(4);
    return true;
}

let result = f2();
// 0
// 1
// Error
// 2
// return false;
console.log(result); // false