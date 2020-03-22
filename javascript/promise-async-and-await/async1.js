~async function hello() {
    try {
        let url = '';
        const resp = await fetch(url);
        console.log(resp);
    } catch (error) {

    }
}();