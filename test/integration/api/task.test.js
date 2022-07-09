/* global beforeAll afterAll describe test expect */
const request = require('supertest')
const config = require('../../../src/config')
const app = require('../../../src/app')
const { sequelize } = require('../../../src/models')
const { signin } = require('../../helpers')
const helper = require('../../helpers/index')

const API_TASK = `${config.API_BASE}/task`

const DEFAULT_TASK = {
    title: 'Todo Task',
    description: 'Todo Task',
    text: 'Todo Task'
}
let USER_TOKEN = ''
beforeAll(async() => {
    USER_TOKEN = await signin()
})

afterAll(async() => {
    await sequelize.close()
})

describe('Test the task path', () => {
    test('It should add new task', async() => {
        const newTask = {
            title: 'Todo Task',
            description: 'Todo Task',
            text: 'Todo Task',
            author: 1
        }
        const response = await request(app)
            .post(API_TASK).send(newTask)
            .set('Cookie', USER_TOKEN)
        expect(response.statusCode).toBe(201)
    })
})