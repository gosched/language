// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes/Class_fields
// Private static fields
// Private instance fields

class User {
    #test

    constructor(id) {
        this.#test = 'test';
        this.id = id;
    }

    show() {
        console.log(this);
        console.log(this.#test);
    }
}

let user1 = new User(1);
user1.show();