const userService = require('../services/services.user');

// POST: /api/public/auth/sign-up
module.exports.signUp = async (ctx) => {
    const user = ctx.request.body;
    const result = await userService.createUser(user);

    return result.success ? ctx.showResult(result.data, result.code) : ctx.showError(result.message, result.code);
}

// POST: /api/public/auth/sign-in
module.exports.signIn = async (ctx) => {
    const { username = '', password = '' } = ctx.request.body;
    const result = await userService.login(username, password);

    return result.success ? ctx.showResult(result.data, result.code) : ctx.showError(result.message, result.code);
}