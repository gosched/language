// closure

// 定義在一個函數內部的函數
// 可以讀取外部函數的變數
// 延長該變數的壽命
// 封裝該變數

function f1() {
    let n = 1000;

    function f2() {
        n--;
        return n;
    }

    return f2;
}

let result = f1();
console.log(result()); // 999
console.log(result()); // 998
