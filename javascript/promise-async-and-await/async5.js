// function request(x) {
//     return new Promise(function (resolve, reject) {
//         x = x ** x
//         if (x % 2 != 0) {
//             resolve(x);
//         } else {
//             reject(new Error('x == even'))
//         }
//     });
// }

// async function hello(x) {
//     let result = await request(x);
//     console.log(result);
//     return result;
// }

// hello(11)
//     .then(function (result) {
//         return result + 1;
//     })
//     .then(function (result) {
//         console.log(result);
//     })
//     .catch(function (error) {
//         console.log(error);
//     })
//     .finally(function () {
//         console.log("finally");
//     })

// ~async function () {
//     console.log('begin');
//     let url = '';
//     await fetch(url)
//         .then(result => {
//             return result.json();
//         })
//         .then(result => {
//             let city = result.city;
//             let temperature = result.temperature;
//             console.log(`city: ${city} temperature:${temperature}`);
//         });
//     console.log('end');
// }();
