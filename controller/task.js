const { User } = require('../model/user')
const userService = require('../service/user')
const taskService = require('../service/task')
const { IN_SUFFICIENT_QUOTA, ERROR } = require('../util/constant')

const taskController = {
    createTask: async (req, res, next) => {
        try {
            const user = await User.findOne({ _id: req.user.id, deleted_at: null });
            if (!user) {
                throw ERROR.USER_NOT_FOUND
            }
            // check user quota
            const { newRemainingPost, createdAt } = userService.checkUserQuota(user.quota)
            if (newRemainingPost === IN_SUFFICIENT_QUOTA) {
                throw ERROR.IN_SUFFICIENT_QUOTA
            }
            // create post
            const task = await taskService.createTask(user, req.body);
            // update user quota
            await userService.updateUserQuota(user, newRemainingPost, createdAt)
            return res.status(201).send({
                status: 'success',
                id: task._id
            })
        } catch (error) {
            next(error)
        }
    }
}

module.exports = taskController