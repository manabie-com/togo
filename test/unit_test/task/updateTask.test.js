const TaskModel = require('../../../src/models/task')
const {
    updateTaskService
} = require('../../../src/services/task')

describe('Task: Update task', () => {
    it("Update task with taskId and createdById correct reference", async () => {
        const data = {
            id: "61bef6594815ad48a1afe950",
            title: "task 1",
            description: "description of task 1",
            createdDate: "2021-12-19T15:49:39.038Z",
            createdById: "61bdeb5edbf6722d48e24d3c"
        }
        TaskModel.findByIdAndUpdate = jest.fn().mockResolvedValue(data)
        let newTask = await updateTaskService(data.id, data)
        expect(newTask).toStrictEqual(data)
    })
    it("Update task with taskID and userID incorrect reference", async () => {
        try {
            const data = {
                id: "61bef6594815ad48a1afe950",
                title: "task 1",
                description: "description of task 1",
                createdDate: "2021-12-19T15:49:39.038Z",
                createdById: "61bdeb5edbf6722d48e24d3c"
            }
          await updateTaskService(data.id, data)
        } catch (err) {
          expect(err).toThrow(TypeError);
        }
      });
})