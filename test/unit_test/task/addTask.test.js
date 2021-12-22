const TaskModel = require('../../../src/models/task')
const {insertTaskService} = require('../../../src/services/task')

describe('Task: Insert task', ()=>{
    it("Check create task with createdById correct reference", async()=>{
        const data ={
            title: "task 1",
            description: "description of task 1",
            createdDate: "2021-12-19T15:49:39.038Z",
            createdById : "61bdeb5edbf6722d48e24d3c"
        }

        TaskModel.create = jest.fn().mockResolvedValue(data)
        let newTask = await insertTaskService(data)
        expect(newTask).toStrictEqual(data)
    })

    it("Check create task with createdById incorrect reference", async()=>{
        try {
            const data ={
                title: "task 1",
                description: "description of task 1",
                createdDate: "2021-12-19T15:49:39.038Z",
                createdById : "61bdeb5edbf6722d48e24d3a"
            }
    
            TaskModel.create = jest.fn().mockResolvedValue(data)
            await insertTaskService(data)
        } catch (error) {
            expect(error).toThrow(TypeError);
        }
        
    })
})