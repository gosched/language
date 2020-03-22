// Variadic function

function sum(...theArgs) {
    let total = 0;
    for (let num of theArgs) {
        total += num;
    }
    return total;
}

console.log(sum(1, 2, 3));

console.log(sum(1, 2, 3, 4));