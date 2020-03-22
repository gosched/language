function getMyFirstPromise(x) {
    return new Promise(function (resolve, reject) {
        x = x ** x
        if (x % 2 != 0) {
            resolve(x);
        } else {
            reject(new Error('x == even'))
        }
    });
}

let number = Math.floor(Math.random() * 10);

let p1 = getMyFirstPromise(number);

p1
    .then(function onfulfilled(result) {
        console.log(result);
    })
    .catch(function onrejected(error) {
        // console.log(error);
    })