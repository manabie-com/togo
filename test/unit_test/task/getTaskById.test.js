const TaskModel = require('../../../src/models/task')
const {
    getTaskByIdService
} = require('../../../src/services/task')
describe('Task: Get task by Id ', () => {
    it("Get task exist", async () => {
        const data = 
            {
            id: "61bef6594815ad48a1afe950",
            title: "task 1",
            description: "description of task 1",
            createdDate: "2021-12-19T15:49:39.038Z",
            createdById: "61bdeb5edbf6722d48e24d3c"
        }
        TaskModel.findById = jest.fn().mockResolvedValue(data)
        let newTask = await getTaskByIdService(data.id)
        expect(newTask).toStrictEqual(data)
    })

    it("Get task not exist", async () => {
        try {
            const id = "61bef6594815ad48a1afe950"
            await getTaskByIdService(id)
        } catch (err) {
          expect(err).toThrow(TypeError);
        }
    })
})