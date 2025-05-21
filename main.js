function getUser() {
    return { name: 'Ashish' };
}

function main() {
    try {
        const user = getUser()
    } catch(err) {
        console.log(err)
    }
}
