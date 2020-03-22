function hello(i, time) {
    return new Promise((resolve, reject) => {
        setTimeout(_ => resolve(i), time);
    })
}

async function foo() {
    console.log("foo");
    let x = await hello(5, 5000);
    let y = await hello(10, 10000);
    console.log('foo done', x, y)
}

foo();