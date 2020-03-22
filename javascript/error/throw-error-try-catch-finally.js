function pow(x, y) {
    if (x === undefined || y === undefined) {
        throw ReferenceError;
        // throw Error;
    }

    if (typeof x !== 'number' || typeof y !== 'number' ) {
        throw TypeError;
    }

    if (y == 0) {
        return 1;
    }

    let z = 1;
    while (y > 0) {
        z *= x;
        y--;
    }

    return z;
}

items = [[0,3], [3,0], [3,3]];

items.map((data) => {
    let result = pow(data[0], data[1]);
    console.log(result);
});

for ([x,y] of items) {
    try {
        let result = pow(x, y);
        console.log(result);
    } catch (error) {
        console.log(error);
    }
}

try {
    throw '[Error: error !!!]';
} catch (error) {
    console.log(error);
}

try {
    pow(1);
} catch (error) {
    console.log(error);
}

try {
    pow(2, '3');
} catch (error) {
    console.log(error);
} finally {
    console.log('finally...');
}

// console.assert(1 < 0, '1 < 0');
