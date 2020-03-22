"use strict";

console.log(this); // {}

function showThis() {
  console.log(this);        // undefined
  console.log(typeof this); // undefined
}

showThis();

function Apple(id) {
  this.id = id;
}

Apple.prototype.show = function () {
  console.log(this);
};

Apple.prototype.showID = function () {
  console.log(this.id);
};

let apple = new Apple(4);
apple.show();
apple.showID();

let x = 1;
let a = {
  x: 2,
  b: {
    x: 3,
    c: function () {
      console.log(this.x)
    }
  }
};

a.b.c(); // 3

a.c = a.b.c;
a.c();  // 2

let c = a.b.c;
// c(); // undefined or TypeError

let o = {
  v: 'hello',
  items: ['i1', 'i2'],
  hello: function f() {
    this.items.forEach(function (item) {
      console.log(this.v + ' ' + item);
    }, this);
  }
}

o.hello()
