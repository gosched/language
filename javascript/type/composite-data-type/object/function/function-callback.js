// callback

// 將一個回調函數作為另一個函數的參數

let sum = function (x, y) {
    return x + y;
};

let product = function (x, y) {
    return x * y;
};

let calculate = function (x, y, callback) {
    console.log(x, y, callback(x, y));
};

calculate(3, 5, sum);
calculate(3, 5, product);