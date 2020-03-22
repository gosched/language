~async function hello() {
    let url = '';

    const resp = await fetch(url).catch(error => {
        return {
            ok: false,
            status: -1,
            error: error,
        };
    });

    if (resp.status === '') {

    } else {
        console.log(resp.data);
    }
}()