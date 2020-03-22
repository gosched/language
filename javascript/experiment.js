// let str = 'zyx';
// let array = str.split('');
// array.sort((a,b)=>{ return a > b ? 1 : -1});
// str = array.join('');

// let MergeSort = function () {

// };

// let QuickSort = function (items) {
//     quickSort(items, 0, items.length - 1);
// };

// let quickSort = function (items, low, high) {
//     if (low < high) {
//         let p = partition(items, low, high);
//         quickSort(low, p - 1);
//         quickSort(p + 1, high);
//     }
// };

// let partition = function (items, low, high) {
//     let pivot = items[high];
//     let index = low;
//     for (let i = low; i < high; i++) {
//         if (items[i] < pivot) {
//             [items[i], items[index]] = [items[index], items[i]];
//             index++;
//         }
//     }
//     [items[high], items[index]] = [items[index], items[high]];
//     return index;
// };

// let items1 = [11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7];
// let items2 = [101, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7];
// console.log(items1);
// console.log(items2);

// items1 = MergeSort(items1);
// items2 = MergeSort(items2);
// console.log(items1);
// console.log(items2);

// items1.reverse();
// items2.reverse();

// QuickSort(items1);
// QuickSort(items2);
// console.log(items1);
// console.log(items2);

// async function async1() {
//     console.log('async1 start');
//     await async2();
//     console.log('async1 end');
// }
// async function async2() {
//     console.log('async2');
// }

// console.log('script start');

// setTimeout(function () {
//     console.log('setTimeout');
// }, 0)

// async1();

// new Promise(function (resolve) {
//     console.log('promise1');
//     resolve();
// }).then(function () {
//     console.log('promise2');
// });

// console.log('script end');

// script start
// async1 start
// async2
// promise1
// script end
// async1 end
// promise2
// setTimeout