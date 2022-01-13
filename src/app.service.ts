import { BadRequestException, Injectable, NotFoundException } from '@nestjs/common'
import { join } from 'path'
import { Low, JSONFile } from 'lowdb'
import { chain } from 'lodash'
import { TakeTaskDTO } from './dto/dto'
import { TTask, TUser } from './typing'

const TASK_FILE_PATH = join('src/database/task.json')
const USER_FILE_PATH = join('src/database/user.json')
@Injectable()
export class AppService {
  private readonly taskDB
  private readonly userDB
  constructor() {
    const taskAdapter = new JSONFile(TASK_FILE_PATH)
    const userAdapter = new JSONFile(USER_FILE_PATH)
    this.taskDB = new Low(taskAdapter)
    this.userDB = new Low(userAdapter)
  }

  async takeTask(userId = 0, input: TakeTaskDTO): Promise<TTask> {
    const { taskId } = input
    await this.validate(userId, taskId)

    const res = await this.assignTask(userId, taskId)
    return res
  }

  async validate(userId = 0, taskId = 0): Promise<void> {
    const [user, userTasksCount, task] = await Promise.all([
      this.getUser(userId),
      this.getUserTasksCount(userId),
      this.getTask(taskId),
    ])

    if (!user) {
      throw new NotFoundException('User not found')
    }

    if (!task) {
      throw new NotFoundException('Invalid task')
    }

    if (userTasksCount >= user.limitTask) {
      throw new BadRequestException('Over limit task per day')
    }

  }

  async getUserTasksCount(userId = 0): Promise<number> {
    await this.taskDB.read()
    this.taskDB.chain = chain(this.taskDB.data)
    const userTasks = this.taskDB.chain
      .filter((task) => {
        const isTakenTask = task.assignedUserId === userId
        const now = new Date().toDateString()
        const assignedDate = new Date(task.assignedDate).toDateString()
        const isToday = assignedDate === now
        return isTakenTask && isToday
      })
      .value()

    return (userTasks || []).length
  }

  async getUser(userId = 0): Promise<TUser> {
    await this.userDB.read()
    this.userDB.chain = chain(this.userDB.data)
    const user = this.userDB.chain
      .find({ id: userId })
      .value()

    return user
  }

  async getTask(taskId = 0): Promise<TUser> {
    await this.taskDB.read()
    this.taskDB.chain = chain(this.taskDB.data)
    const user = this.taskDB.chain
      .find({
        id: taskId,
        assignedUserId: 0,
      })
      .value()
    return user
  }

  async assignTask(userId = 0, taskId = 0): Promise<TTask> {
    await this.taskDB.read()
    const assignTask = this.taskDB.data.reduce((acc, task) => {
      if (task.id != taskId) {
        return acc
      }

      task.assignedUserId = userId
      task.assignedDate = Date.now()
      acc = task
      return acc
    }, {})
    this.taskDB.write()

    return assignTask
  }
}
