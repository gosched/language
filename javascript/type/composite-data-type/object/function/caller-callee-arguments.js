function f1(a1, a2, a3) {
    console.log('f1');
    console.log(arguments);
    console.log(arguments.callee);
    console.log(f1.caller);
    f2();
}

function f2() {
    console.log('f2');
    console.log(arguments);
    console.log(arguments.callee);
    console.log(f2.caller);
}

f1(1, 2, 3);