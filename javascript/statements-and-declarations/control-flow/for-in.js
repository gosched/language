// The for...in statement iterates over all non-Symbol, enumerable properties of an object.

let object1 = { 'a': 1, 'b': 2, 'c': 3 };

for (let property in object1) {
    console.log(property);
}

// a
// b
// c