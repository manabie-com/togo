const User = require('../models/user.model')
const UserSetting = require('../models/userSetting.model')
const {InvalidUser} = require("../utils/errors");

module.exports = {
    getAllUsers: async () => {
        return User.find({}, {id: 1, username: 1, created: 1}, {});
    },
    getUser: async (userId) => {
        try {
            return await User.findOne({_id: userId}, {}, {})
        } catch (err) {
            throw new InvalidUser(`not found userId: ${userId}`)
        }
    },
    getUserSetting: async (userId) => {
        try {
            return await UserSetting.findOne({user: userId}, {}, {})
        } catch (err) {
            throw new InvalidUser(`not found userId: ${userId}`)
        }
    }
}