const UserRepository = require('../repositories/user.repository');

exports.getUsers = (req, res, next) => {
    UserRepository.getUsers(req, res, next);
}


