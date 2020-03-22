class Person {
    constructor(name) {
        console.log("Person constructor");
        this.name = name;
    }

    sayHello() {
        console.log(`Hello, I am ${this.name}!`);
    }
}

class User extends Person {
    constructor(id, name) {
        super(name); // invoke constructor of base class.
        this.id = id;
    }

    sayHello() { // method overriding
        console.log(`Hello World, I am ${this.name}!`);
    }
}

let p1 = new Person('apple');
console.log(p1);
p1.sayHello();

let u1 = new User(2, 'github');
console.log(u1);
u1.sayHello();