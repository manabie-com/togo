const dbConfig = require('../config/database')
const Sequelize = require('sequelize')
const sequelize = new Sequelize(dbConfig)
const moment = require('moment')
class TaskService {
    constructor(Task) {
        this.task = Task
    }

    async create(taskDTO) {
        try {
            const verifyLimitTask = await this.verifyTaskLimit(taskDTO.author)
            if (verifyLimitTask) {
                const task = await this.task.create(taskDTO)
                return { task, status: 201 }
            } else {
                throw new Error('create task fail limit task for day')
            }
        } catch (err) {
            throw new Error(err.message)
        }
    }

    async verifyTaskLimit(userId) {
        const newDay = new Date()
        const fromDate = moment(newDay).format("YYYY-MM-DD")
        const toDate = moment(fromDate).add(1, 'days').format("YYYY-MM-DD")
        const query = `SELECT * FROM "task" where author = ${userId} and "createdAt" between '${fromDate}' and '${toDate}'`;
        const countTask = await sequelize.query(query)
        const queryUser = `SELECT * FROM "users" where id = ${userId}`
        const user = await sequelize.query(queryUser)
        const queryConfig = `SELECT * FROM "configs" where role = '${user[0][0].role}'`
        const config = await sequelize.query(queryConfig)
        if (config[0][0].limit <= countTask[0].length) {
            return false
        }
        return true
    }
}

module.exports = TaskService