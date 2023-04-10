const { Task } = require('../model/task');

const taskService = {
    createTask: async (user, body) => {
        const task = await Task.create({
            author: user._id,
            ...body
        })
        return task;
    }
}

module.exports = taskService