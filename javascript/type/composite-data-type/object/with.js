// 操作同一個對象的多個屬性

let o1 = {
    a: 1,
    b: 2,
    c: 3
};

// with 沒有改變作用域
with (o1) {
    a = 101;
    b = 102;
    c = 103;
    test = 123; // 會創造一個當前作用域的全局變量 test
    console.log(o1); // { a: 101, b: 102, c: 103 }
}

console.log(o1); // { a: 101, b: 102, c: 103 }