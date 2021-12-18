const knex = require('../db/knex').knex;
const { userSchema } = require('../schema');
const { EMAIL_DUPLICATE_ERROR } = require('../config');

// Check to see if there is already a user with the same email
const validateUser = async (req, res, next) => {
    // Check body validity
    const { error } = userSchema.validate(req.body);
    if (error) {
        res.status(400).json(error);
    }

    // look for existed email
    const results = await knex('users').where('email', req.body.email).select('*');
    if (results.length === 0) {
        next();
    } else {
        res.status(400).json({ message: EMAIL_DUPLICATE_ERROR });
    }
}

module.exports = {
    validateUser
}
