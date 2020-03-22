// 不像 goroutine
// 會執行完再退出程序
setTimeout(function () {
    console.log('timeout')
}, 1000)

// function timeout(ms) {
//     return new Promise((resolve, reject) => {
//         setTimeout(resolve, ms, 'done');
//     });
// }

// timeout(100)
//     .then((value) => {
//         console.log(value);
//     })

// function timeoutWithReturn() {
//     return new Promise((resolve) => {
//         console.log('timeoutWithReturn...');
//         setTimeout(() => {
//             console.log('timeout!!!');
//             return resolve('resolve value');
//         }, 3000);
//     });
// }

// (async function () {
//     let result = await timeoutWithReturn();
//     console.log(result);
// })();
