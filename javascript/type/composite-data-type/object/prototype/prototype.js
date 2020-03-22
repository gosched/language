// let object1 = new Constructor();
// let object2 = new Constructor();

// 透過 constructor 生成的 object1 與 object2
// object1 的 property 與 object2 的 property 是獨立的

// 透過 constructor 的 prototype 屬性
// 來 同時設置 同時調整 需要共享的 property

function User(id) {
    this.id = id;
}

User.prototype.hello = function () {
    console.log("Hello", this.id);
};

let user1 = new User(1);
let user2 = new User(2);

console.log(User.prototype);
console.log(user1);
console.log(user2);

user1.hello();
user2.hello();

function f1() {

};

let f2 = function() {

};

let f3 = () => {
    
};

let o1 = {};

console.log(f1.prototype); // f1 {}
console.log(f2.prototype); // f2 {}
console.log(f3.prototype); // undefined
console.log(o1.prototype); // undefined
