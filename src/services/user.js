const User = require('../models/user')
const TaskLimit = require('../models/taskLimit')

const createUserService = async (data)=>{
    try {
        const checkUserExist = await User.findOne({ email: data.email} )
        if (checkUserExist) {
            throw new Error('Email is used')
        }
        let user = new User(data);
        user = await User.create(user);
        const token = await user.generateAuthToken();
        return {user, token}
    } catch (error) {
        throw error
    }
}

const loginService = async(email, password)=>{
    try {
        const user = await User.findByCredentials(email, password)
        if (!user) {
            return null;
        }
        const token = await user.generateAuthToken();
        return {user, token};
    } catch (error) {
        throw error;
    }
}

module.exports = {createUserService, loginService}