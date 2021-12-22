const TaskModel = require('../../../src/models/task')
const UserModel = require('../../../src/models/user')
const {getTaskByUserIdService} = require('../../../src/services/task')
const moment = require('moment')

describe('Task: Get tasks by user id', () => {
    it("Get task with user exist", async () => {
        const userId = "61bdeb5edbf6722d48e24d3c"
        const data = [
            {
                id: "61bef6594815ad48a1afe950",
                title: "task 1",
                description: "description of task 1",
                createdDate: "2021-12-19T15:49:39.038Z",
                createdById: "61bdeb5edbf6722d48e24d3c"
            },
            {
                id: "61bef6594815ad48a1afe952",
                title: "task 2",
                description: "description of task 2",
                createdDate: "2021-12-19T15:49:39.038Z",
                createdById: "61bdeb5edbf6722d48e24d3c"
            }
        ]
        UserModel.findById = jest.fn().mockResolvedValue({})
        TaskModel.find = jest.fn().mockResolvedValue(data)
        let newTask = await getTaskByUserIdService(userId, moment("2021-12-21"))
        expect(newTask).toStrictEqual(data)
    })

    it("Get task with user not exist", async () => {
        try {
            const userId = "61bdeb5edbf6722d48e24d3c"
            await getTaskByUserIdService(userId, moment("2021-12-21"))
        } catch(err) {
            expect(err).toThrow(TypeError)
        }
    })
})