# Create

## Reference

https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/create

## Example 1

```javascript
function User(id = 0) {
    this.id = id;
}

User.prototype.showID = function () {
    console.log(this.id);
};

let u1 = new User(1);
console.log(u1);
u1.showID();
```

## Example 2

```javascript
let language = {
    name: '',
    show: function () {
        console.log(`name:${this.name}`);
    }
};

let javascript = Object.create(language);
console.log(javascript);        // {}
javascript.name = 'JavaScript';
console.log(javascript);        // { name: 'JavaScript' }
javascript.show();
```

## Example 3

```javascript
class Person {
    constructor(id) {
        this.id = id;
    }
    showID() {
        console.log(this.id);
    }
}

let p1 = new Person(1);
console.log(p1);
p1.showID();
```