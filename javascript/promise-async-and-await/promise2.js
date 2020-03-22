const delay = (s) => {
    return new Promise(resolve => {
        setTimeout(resolve, s);
    });
};

delay()
    .then(() => {
        console.log(1);
        return delay(1000);
    })
    .then(() => {
        console.log(2);
        return delay(2000);
    })
    .then(() => {
        console.log(3);
    });

~async function () {
    console.log(1);
    await delay(1000);
    console.log(2);
    await delay(2000);
    console.log(3);
}();

(async function () {
    console.log(1);
    await delay(1000);
    console.log(2);
    await delay(2000);
    console.log(3);
})();