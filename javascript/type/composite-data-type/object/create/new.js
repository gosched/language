function Apple() {
    if (!new.target) {
        throw new Error('need new');
    }
}

let a1 = Apple();
let a2 = new Apple();