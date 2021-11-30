const checkAuth = (username, password) => {
    if (username && password) {
        return username === "firstUser" && password === "example";
    } else {
        return false;
    }
};

module.exports = {
    checkAuth,
};