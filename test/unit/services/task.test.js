/* global jest beforeAll test expect */
const TaskService = require('../../../src/services/task')
jest.mock('../../../src/models')
const { Task } = require('../../../src/models')

let taskService
beforeAll(() => {
    taskService = new TaskService(Task)
})

test('should send task', async() => {
    const newTask = {
        title: 'Todo Task',
        description: 'Todo Task',
        text: 'Todo Task',
        author: 1
    }
    Task.create.mockResolvedValue(newTask)
    const resp = await taskService.create(newTask)
    expect(resp.status).toBe(201);
})