function Person(name) {
    this.name = name;
}

Person.prototype.skill = 'coding';

Person.prototype.sayHello = function () {
    console.log('hello');
};

function User(name) {
    this.name = name;
}

User.prototype = new Person();

let user = new User();
console.log(user.skill)
console.log(user instanceof Person);
