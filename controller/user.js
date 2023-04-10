const { User } = require('../model/user')
const userService = require('../service/user')
const jwtService = require('../util/jwt')
const { ERROR } = require('../util/constant')

const userController = {
    getUser: async (req, res, next) => {
        try {
            // TODO: validate input
            const user = await User.findOne({ _id: req.params.id, deleted_at: null }, { email: true, quota: true }).lean();
            if (!user) {
                throw ERROR.USER_NOT_FOUND
            }
            res.status(200).send(user)
        } catch (error) {
            next(error)
        }
    },
    register: async (req, res, next) => {
        try {
            // Validate input
            if (!req.body.email || !req.body.password) {
                throw ERROR.INVALID_INPUT
            }
            const userExisted = await User.findOne({ email: req.body.email });
            if (userExisted) {
                throw ERROR.USER_EXISTED
            }
            // TODO: for simplicity of demo, I dont encrypt the password
            const user = await userService.createUser(req.body)
            return res.status(201).send({
                status: 'success',
                _id: user._id
            })
        } catch (error) {
            next(error)
        }
    },
    login: async (req, res, next) => {
        try {
            // TODO: validate input
            const user = await User.findOne({ email: req.body.email, deleted_at: null });
            if (!user) {
                throw ERROR.USER_NOT_FOUND
            }
            // TODO: for simplicity of demo, I dont encrypt the password
            if (user.password !== req.body.password) {
                throw ERROR.INVALID_CREDENTIAL
            }
            res.status(200).send({
                _id: user._id,
                access_token: jwtService.generateJwtToken({ id: user._id }),
            })
        } catch (error) {
            next(error)
        }
    }
}

module.exports = userController