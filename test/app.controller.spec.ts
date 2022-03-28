import { ForbiddenException, NotFoundException } from '@nestjs/common'
import { AppController } from '../src/app.controller'
import { AppService } from '../src/app.service'
import { Low, JSONFile } from 'lowdb'

import { TASK_FILE_PATH } from '../src/utils/constant'

const taskAdapter = new JSONFile(TASK_FILE_PATH)
const taskDB = new Low(taskAdapter)


// TODO add recover data after test
describe('taskController', () => {
  let taskService: AppService
  let taskController: AppController

  beforeEach(async () => {
    taskService = new AppService()
    taskController = new AppController(taskService)
  })

  describe('takeTask', () => {
    it('should return task taken info', async () => {
      await taskDB.read()
      const backupData = taskDB.data
      const expectedResult = {
        id: 1,
        name: 'task01',
        isFinished: 0,
        assignedUserId: 2,
      }
      const userId = 2
      const input = { taskId: 1 }
      expect(await taskController.takeTask(userId, input))
        .toMatchObject(expectedResult)
      taskDB.data = backupData
      await taskDB.write()
    })

    it('should throw error user not found', async () => {
      const userId = 0
      const input = { taskId: 1 }
      await expect(taskController.takeTask(userId, input))
        .rejects
        .toThrowError(NotFoundException)
    })

    it('should throw error invalid task', async () => {
      const userId = 2
      const input = { taskId: 20 }
      await expect(taskController.takeTask(userId, input))
        .rejects
        .toThrowError(NotFoundException)
    })

    it('should throw error over limit task', async () => {
      await taskDB.read()
      const backupData = taskDB.data

      const userId = 1
      const validInput = { taskId: 4 }
      await taskController.takeTask(userId, validInput)
      const input = { taskId: 5 }
      await expect(taskController.takeTask(userId, input))
        .rejects
        .toThrowError(ForbiddenException)

      taskDB.data = backupData
      await taskDB.write()
    })

  })
})
