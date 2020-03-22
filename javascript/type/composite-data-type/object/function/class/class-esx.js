// https://en.wikipedia.org/wiki/Object-oriented_programming
// https://en.wikipedia.org/wiki/Prototype-based_programming

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Details_of_the_Object_Model
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes/Class_fields

// Overview
// class == function == syntactic sugar
// instance == object

// Keyword
// class
// constructor
// get, set
// extends, super

// Static fields
// static

// Private fields
// static #PRIVATE_STATIC_FIELD
// #privateField

// 繼承     (inheritance)
// 封裝     (encapsulation) 
// 動態連結  (dynamic binding)

class User {
    static test() {
        console.log('static');
    }

    constructor(id) {
        this._id = id;
    }

    get id() {
        return this._id;
    }

    set username(username) {
        this._username = username;
    }

    get username() {
        return this._username;
    }

    sayHello() {
        console.log(`Hello, I am ${this._username}!`);
    }
}

console.log(User.toString());
console.log(User.prototype);                       // User {}
console.log(User.name);                            // User
console.log(User.length);                          // 1

User.test();

console.log(User == User.prototype.constructor);   // true
console.log(typeof User);                          // function
console.log(Object.prototype.toString.call(User)); // [object Function]

let u1 = new User(1);
console.log(u1.id);
u1.username = 'apple';
console.log(u1.username);
console.log(u1);
u1.sayHello();