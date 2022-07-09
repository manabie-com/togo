/* global jest beforeAll test expect */
const TaskService = require('../../../src/services/task')
jest.mock('../../../src/models')
const { Task } = require('../../../src/models')
const dbConfig = require('../../../src/config/database')
const Sequelize = require('sequelize')
const sequelize = new Sequelize(dbConfig)
const { signin } = require('../../helpers')
let taskService

beforeAll(async() => {
    taskService = new TaskService(Task)
})

test('should send task', async() => {
    const query = `SELECT * FROM "users" where email like 'user@test.com'`
    const user = await sequelize.query(query)
    if (user[0].length <= 0) {
        await signin()
    }
    const newTask = {
        title: 'Todo Task',
        description: 'Todo Task',
        text: 'Todo Task',
        author: user[0][0].id
    }
    const queryConfig = `SELECT * FROM "configs" where role = '${user[0][0].role}'`
    const config = await sequelize.query(queryConfig)
    Task.create.mockResolvedValue(newTask)
    const resp = await taskService.create(newTask)
    expect(resp.status).toBe(201);
})