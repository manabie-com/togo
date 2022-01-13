import { BadRequestException, NotFoundException } from '@nestjs/common'
import { AppController } from './app.controller'
import { AppService } from './app.service'

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
      const result = {
        id: 1,
        name: 'task01',
        isFinished: 0,
        assignedUserId: 2,
      }
      const userId = 2
      const input = { taskId: 1 }
      expect(await taskController.takeTask(userId, input))
        .toMatchObject(result)
    })

    it('should throw error user not found', async () => {
      const userId = 0
      const input = { taskId: 1 }
      await expect(taskController.takeTask(userId, input))
        .rejects
        .toThrowError(NotFoundException)
    })

    it('should throw error invalid user', async () => {
      const userId = 2
      const input = { taskId: 20 }
      await expect(taskController.takeTask(userId, input))
        .rejects
        .toThrowError(NotFoundException)
    })
  })
})
