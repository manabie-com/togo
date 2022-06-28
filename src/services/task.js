const ConfigService = require('../services/config')
const UserService = require('../services/users')
const { User } = require('../models')
const { Config } = require('../models')
const configService = new ConfigService(Config)
const userService = new UserService(User)

class TaskService {
    constructor(Task) {
        this.task = Task
    }

    async create(taskDTO) {
        try {
            const verifyLimitTask = await this.verifyTaskLimit(taskDTO.author)
            if (verifyLimitTask) {
                await this.task.create(taskDTO)
            } else {
                throw new Error('create task fail limit task')
            }
        } catch (err) {
            throw new Error(err.message)
        }
    }

    async verifyTaskLimit(userId) {
        const countTask = await this.task.findAll({
            where: {
                author: userId
            }
        })
        const user = await userService.getById(userId)
        const config = await configService.getLimitByRole(user.role)
        if (config.limit <= countTask.length) {
            return false
        }
        return true
    }
}

module.exports = TaskService