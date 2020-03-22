// https://wangdoc.com/javascript/oop/index.html

// This constructor function may be converted to a class declaration.
let Editor = function (name, version) {
    this.name = name;
    this.version = version;
}

// 透過 constructor function 建構 object 必須使用 new 命令
// 沒有使用 new 命令, this == global, vscode == undefined, name == global variable
let vscode = new Editor('vscode', '1.36.1');
console.log(vscode);
console.log(vscode.name);

// 加上 'use strict';
// 忘了使用 new 命令 直接調用 constructor 會報錯

// 函數內部可以使用 new.target 屬性, 判斷函數調用的時候, 是否使用 new 命令
function f() {
    console.log(new.target === f);
    console.log(new.target === undefined);
}

new f();
f();

// 有時拿不到 constructor function
// 只能拿到一個 object
// 這時可以使用 Object.create() 方法

let person1 = {
    name: 'foo',
    age: 50,
    greeting: function () {
        console.log('Hi! I\'m ' + this.name + '.');
    }
};

let person2 = Object.create(person1);
console.log(person1);
console.log(person2);
console.log(person2.name);
person2.greeting();