function hello(i, time) {
    return new Promise((resolve, reject) => {
        setTimeout(_ => resolve(i), time);
    })
}

const test = async function () {
    console.log("test");
    const promises = [hello(5, 5000), hello(10, 10000)]
    const [x, y] = await Promise.all(promises)
    console.log('test done', [x, y])
}

test();