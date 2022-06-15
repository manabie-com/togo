const User = require('../models/user.model');

exports.getUsers = (req, res, next) => {
    User.find({}).then(user => {
        res.json(user);
    });
}


