const User = require('../models/user');

module.exports = async function auth(req, res, next) {
    const { userId } = req.session;
    // console.log("userId: ", userId);
    res.locals.currentUser = null;
    if (userId) {
        const user = await User.findById(userId)
        if (user) {
            req.currentUser = user;
            res.locals.currentUser = user;
        }
        next();
    } else {
        next();
    }

};